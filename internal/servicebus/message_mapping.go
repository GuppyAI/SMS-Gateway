package servicebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"sms-gateway/internal/messaging"
)

func toServiceBusMessage(message messaging.Message) *azservicebus.Message {
	body := message.GetMessageBody()
	address := message.GetAddress().String()

	return &azservicebus.Message{
		Body: []byte(body),
		ApplicationProperties: map[string]any{
			"address": address,
		},
	}
}

func toInternalMessage(message *azservicebus.ReceivedMessage) (*messaging.Message, error) {
	body := string(message.Body)

	addressProperty := message.ApplicationProperties["address"]
	if addressProperty == nil {
		return nil, messaging.ErrInvalidAddressFormat
	}

	address, err := messaging.ParseAddress(addressProperty.(string))
	if err != nil {
		return nil, err
	}

	return messaging.NewMessage(messaging.Response, *address, body), nil
}
