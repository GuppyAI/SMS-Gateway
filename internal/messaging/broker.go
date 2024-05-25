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
	broker.channels[channel.GetSupportedSchema()] = channel
	broker.lock.Unlock()
}

// Publish publishes messages
func (broker *brokerImpl) Publish(message Message) {
	currentLogger := log.With().
		Str("kind", message.kind.String()).
		Str("address", message.GetAddress().String()).
		Str("message", message.GetMessageBody()).
		Logger()

	currentLogger.Trace().Msgf("Publishing message from address %s.", message.address.String())

	if !broker.allowListCheck(message.GetAddress()) {
		currentLogger.Warn().Msg("Discarded message because of failing allow list check!")

		return
	}

	switch message.kind {
	case Request:
		broker.lock.RLock()
		handler := broker.handler
		broker.lock.RUnlock()

		go (func(handler MessageHandler) {
			if err := handler.Handle(message); err != nil {
				currentLogger.Error().
					Err(err).
					Msgf("Error occurred when processing message")
			}
		})(handler)
	case Response:
		broker.lock.RLock()
		channels := broker.channels[message.GetAddress().GetSchema()]
		broker.lock.RUnlock()

		go (func(channel MessageChannel) {
			if err := channel.SendMessage(message); err != nil {
				currentLogger.Error().
					Err(err).
					Msgf("Error occurred when processing message")
			}
		})(channels)
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
