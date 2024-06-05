package servicebus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/avast/retry-go/v4"
	"github.com/rs/zerolog/log"
	"slices"
	"sms-gateway/internal/messaging"
)

type Listener struct {
	receiver Receiver
}

func NewListener(receiver Receiver) *Listener {
	return &Listener{receiver}
}

// Listen will listen for new messages on the service bus and handle them by converting them to messaging.Message and
// calling the given callback function.
// Messages of unsupported schema will be marked as abandoned by this gateway.
// If an error occurs while retrieving, converting or handling a message, it will retry three times and then return the
// error.
func (listener *Listener) Listen(broker messaging.Broker) error {
	for {
		var messages []*azservicebus.ReceivedMessage

		err := retry.Do(func() error {
			var err error

			messages, err = listener.receiver.ReceiveMessages(context.TODO(), 1, nil)
			if err != nil {
				log.Warn().Err(err).Msg("Could not receive messages from service bus!")
				return err
			}

			return nil
		})
		if err != nil {
			log.Warn().Err(err).Msg("Could not receive messages from service bus after three attempts!")
			return err
		}

		for _, message := range messages {
			err := retry.Do(func() error {
				if err := listener.handleMessage(message, broker); err != nil {
					return err
				}

				return nil
			}, retry.Attempts(3))

			if err != nil {
				log.Warn().Err(err).Msg("Could not handle message after three attempts!")
				return err
			}
		}
	}
}

// handleMessage converts the azservicebus.ReceivedMessage to a messaging.Message,
// checks if it is supported by the given messaging.Broker and publishes it on the given messaging.Broker.
// The azservicebus.ReceivedMessage will be handed back to the queue as abandoned if its schema is not supported by this gateway.
// If there is a conversion error, the azservicebus.ReceivedMessage will be sent to the dead letter queue.
func (listener *Listener) handleMessage(message *azservicebus.ReceivedMessage, broker messaging.Broker) error {
	logger := log.With().Str("message_id", message.MessageID).Logger()

	responseMessage, err := toInternalMessage(message)
	if err != nil {
		// It is assumed that this error only occurs when the message retrieved from the service bus is faulty.
		// Therefore, it will be sent to the dead letter queue and this method does not return an error.
		logger.Error().Err(err).Msg("Could not convert service bus message to internal message!")

		if deadLetterErr := listener.receiver.DeadLetterMessage(context.TODO(), message, nil); deadLetterErr != nil {
			logger.Error().Err(deadLetterErr).Msg("Could not dead letter message after message conversion failed!")
			return deadLetterErr
		}

		return nil
	}

	if !slices.Contains(broker.GetSupportedSchemas(), responseMessage.GetAddress().GetSchema()) {
		// Received message is not supported by this gateway. It will be abandoned for other gateway to possibly handle
		// it.
		logger.Debug().Msg("Received message with unsupported schema!")

		if abandonErr := listener.receiver.AbandonMessage(context.TODO(), message, nil); err != nil {
			logger.Error().Err(abandonErr).Msg("Could not abandon message after message was found to be of an unsupported schema!")
			return abandonErr
		}

		return nil
	}

	if err := listener.receiver.CompleteMessage(context.TODO(), message, nil); err != nil {
		logger.Error().Err(err).Msg("Could not mark message as completed")
		return err
	}

	// The message is published after the received service bus message has been marked as completed to prevent messages
	// from being published twice by this gateway.

	broker.Publish(*responseMessage)

	return nil
}
