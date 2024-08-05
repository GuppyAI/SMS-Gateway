package servicebus

import (
	"context"
	"sms-gateway/internal/messaging"
)

type MessageHandler struct {
	sender Sender
}

func NewMessageHandler(sender Sender) (*MessageHandler, error) {
	return &MessageHandler{sender}, nil
}

func (m MessageHandler) Handle(message messaging.Message) error {
	err := m.sender.SendMessage(context.TODO(), toServiceBusMessage(message), nil)
	if err != nil {
		return err
	}

	return nil
}
