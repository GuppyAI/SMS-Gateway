package service_bus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"sms-gateway/internal/messaging"
)

type MessageHandler struct {
	sender *azservicebus.Sender
}

func NewMessageHandler(sender *azservicebus.Sender) (*MessageHandler, error) {
	return &MessageHandler{sender}, nil
}

func (m MessageHandler) Handle(message messaging.Message) error {
	err := m.sender.SendMessage(context.TODO(), toServiceBusMessage(message), nil)
	if err != nil {
		return err
	}

	return nil
}
