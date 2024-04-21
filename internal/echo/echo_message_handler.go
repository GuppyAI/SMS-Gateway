package echo

import (
	"sms-gateway/internal/application_context"
	"sms-gateway/internal/messaging"
)

type MessageHandler struct{}

func (e MessageHandler) Handle(message messaging.Message) error {
	broker := application_context.GetContext().GetBroker()

	newMessage := messaging.NewMessage(messaging.Response, message.GetAddress(), message.GetMessageBody())
	broker.Publish(*newMessage)

	return nil
}

func New() *MessageHandler {
	return &MessageHandler{}
}
