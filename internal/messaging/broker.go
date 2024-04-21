package messaging

import (
	"github.com/rs/zerolog/log"
	"sms-gateway/internal/configuration"
	"sync"
)

type Broker interface {
	AddMessageChannel(MessageChannel)
	Publish(Message)
}

// brokerImpl is used to interact with the pubsub architecture
type brokerImpl struct {
	channels map[AddressSchema]MessageChannel
	handler  MessageHandler
	lock     sync.RWMutex
}

// NewBroker constructs a new brokerImpl
func NewBroker(handler MessageHandler) Broker {
	return &brokerImpl{
		channels: map[AddressSchema]MessageChannel{},
		handler:  handler,
	}
}

// AddMessageChannel adds a new message channel to the message channel pool
func (broker *brokerImpl) AddMessageChannel(channel MessageChannel) {
	broker.lock.Lock()
	defer broker.lock.Unlock()

	broker.channels[channel.GetSupportedSchema()] = channel

	err := channel.ReceiveMessages(broker)
	if err != nil {
		log.Error().Err(err).Msg("Error occurred while receiving messages")
	}
}

// Publish publishes messages
func (broker *brokerImpl) Publish(message Message) {
	log.Trace().Msgf("Publishing message from address %s.", message.address.String())

	if !broker.allowListCheck(message.GetAddress()) {
		log.Warn().
			Str("address", message.GetAddress().String()).
			Str("message", message.GetMessageBody()).
			Msg("Discarded message because of failing allow list check!")

		return
	}

	switch message.kind {
	case Request:
		go (func(handler MessageHandler) {
			if err := handler.Handle(message); err != nil {
				log.Error().
					Err(err).
					Msgf(
						"Error occurred when processing message of type %s with address %s",
						message.kind,
						message.address.String(),
					)
			}
		})(broker.handler)
	case Response:
		go (func(channel MessageChannel) {
			if err := channel.SendMessage(message); err != nil {
				log.Error().
					Err(err).
					Msgf(
						"Error occurred when processing message of type %s with address %s",
						message.kind,
						message.address.String(),
					)
			}
		})(broker.channels[message.GetAddress().GetSchema()])
	}

}

func (broker *brokerImpl) allowListCheck(address Address) bool {
	config := configuration.GetConfig()
	allowList := config.Strings("messaging.allowlist")

	if len(allowList) == 0 {
		return true
	}

	for _, entry := range allowList {
		if address.String() == entry {
			return true
		}
	}

	return false
}
