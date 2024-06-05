package servicebus

import (
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/stretchr/testify/assert"
	"sms-gateway/internal/messaging"
	"testing"
)

func TestToInternalMessage(t *testing.T) {
	serviceBusMessage := &azservicebus.ReceivedMessage{
		ApplicationProperties: map[string]any{
			"address": "sms://+49123456789",
		},
		Body: []byte("This is a test message"),
	}

	internalMessage, err := toInternalMessage(serviceBusMessage)
	if err != nil {
		t.Error(err)
		return
	}

	assert.Equal(t, "This is a test message", internalMessage.GetMessageBody())
	assert.Equal(t, "+49123456789", internalMessage.GetAddress().GetAddress())
	assert.Equal(t, messaging.SMS, internalMessage.GetAddress().GetSchema())
}

func TestToInternalMessage_InvalidFormat(t *testing.T) {
	type TestParams struct {
		Name    string
		Address any
	}

	for _, testParams := range []TestParams{
		{
			Name:    "Invalid address schema",
			Address: "invalid://SomeAddress",
		},
		{
			Name:    "No address schema",
			Address: "+49123456789",
		},
		{
			Name:    "Whitespace address",
			Address: "",
		},
		{
			Name:    "Nil address schema",
			Address: nil,
		},
	} {
		t.Run(testParams.Name, func(t *testing.T) {
			serviceBusMessage := &azservicebus.ReceivedMessage{
				ApplicationProperties: map[string]any{
					"address": testParams.Address,
				},
				Body: []byte("This is a test message!"),
			}

			internalMessage, err := toInternalMessage(serviceBusMessage)

			assert.Nil(t, internalMessage)
			assert.NotNil(t, err)
			assert.ErrorIs(t, messaging.ErrInvalidAddressFormat, err)
		})
	}
}

func TestToServiceBusMessage(t *testing.T) {
	internalMessage := messaging.NewMessage(messaging.Response, *messaging.NewAddress(messaging.SMS, "+49123456789"), "This is a test message")

	serviceBusMessage := toServiceBusMessage(*internalMessage)

	assert.Equal(t, "sms://+49123456789", serviceBusMessage.ApplicationProperties["address"].(string))
	assert.Equal(t, "This is a test message", string(serviceBusMessage.Body))
}
