package echo

import (
	"sms-gateway/internal/messaging"
)

type MessageHandler struct {
	broker messaging.Broker
}

func (e MessageHandler) Handle(message messaging.Message) error {
	newMessage := messaging.NewMessage(messaging.Response, message.GetAddress(), message.GetMessageBody())
	e.broker.Publish(*newMessage)

	return nil
}

func New(broker messaging.Broker) *MessageHandler {
	return &MessageHandler{broker}
}
