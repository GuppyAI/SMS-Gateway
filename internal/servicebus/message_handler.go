package servicebus

import (
	"context"
	"github.com/rs/zerolog/log"
	"sms-gateway/internal/messaging"
)

type MessageHandler struct {
	sender Sender
}

func NewMessageHandler(sender Sender) (*MessageHandler, error) {
	return &MessageHandler{sender}, nil
}

func (m MessageHandler) Handle(message messaging.Message) error {
	log.Debug().
		Str("address", message.GetAddress().String()).
		Str("message", message.GetMessageBody()).
		Msg("Sending message to service bus!")

	err := m.sender.SendMessage(context.TODO(), toServiceBusMessage(message), nil)
	if err != nil {
		return err
	}

	return nil
}
