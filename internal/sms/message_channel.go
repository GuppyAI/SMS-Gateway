package sms

import (
	"github.com/rs/zerolog/log"
	"github.com/warthog618/modem/gsm"
	"sms-gateway/internal/messaging"
)

type MessageChannel struct {
	sender   Sender
	receiver Receiver
}

func NewMessageChannel(sender Sender, receiver Receiver) messaging.MessageChannel {
	return &MessageChannel{sender: sender, receiver: receiver}
}

func (m MessageChannel) GetSupportedSchema() messaging.AddressSchema {
	return messaging.SMS
}

func (m MessageChannel) SendMessage(message messaging.Message) error {
	if message.GetAddress().GetSchema() != m.GetSupportedSchema() {
		return messaging.UnsupportedSchemaErr
	}

	log.Debug().
		Str("phone_number", message.GetAddress().GetAddress()).
		Str("message", message.GetMessageBody()).
		Msg("Sending SMS message!")

	err := m.sender.SendSMS(message.GetAddress().GetAddress(), message.GetMessageBody())
	if err != nil {
		return err
	}

	return nil
}

func (m MessageChannel) ReceiveMessages(broker messaging.Broker) error {
	m.receiver.Listen(func(message gsm.Message) {
		address := messaging.NewAddress(messaging.SMS, message.Number)

		brokerMessage := messaging.NewMessage(messaging.Request, *address, message.Message)

		broker.Publish(*brokerMessage)
	})

	return nil
}
