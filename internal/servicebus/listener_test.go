package servicebus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/messaging"
	"testing"
)

func TestListener_HandleMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	receiver := NewMockReceiver(ctrl)
	listener := NewListener(receiver)
	broker := messaging.NewMockBroker(ctrl)

	serviceBusMessage := &azservicebus.ReceivedMessage{
		Body: []byte("This is a test message"),
		ApplicationProperties: map[string]any{
			"address": "sms://+49123456789",
		},
	}

	message := messaging.NewMessage(messaging.Response, *messaging.NewAddress(messaging.SMS, "+49123456789"), "This is a test message")

	broker.EXPECT().GetSupportedSchemas().Return([]messaging.AddressSchema{messaging.SMS})
	receiver.EXPECT().CompleteMessage(context.TODO(), serviceBusMessage, nil)
	broker.EXPECT().Publish(*message)

	err := listener.handleMessage(serviceBusMessage, broker)
	if err != nil {
		t.Error(err)
	}
}

func TestListener_HandleMessage_UnsupportedSchema(t *testing.T) {
	ctrl := gomock.NewController(t)
	receiver := NewMockReceiver(ctrl)
	listener := NewListener(receiver)
	broker := messaging.NewMockBroker(ctrl)

	serviceBusMessage := &azservicebus.ReceivedMessage{
		Body: []byte("This is a test message"),
		ApplicationProperties: map[string]any{
			"address": "email://+49123456789",
		},
	}

	broker.EXPECT().GetSupportedSchemas().Return([]messaging.AddressSchema{})
	receiver.EXPECT().AbandonMessage(context.TODO(), serviceBusMessage, nil)

	err := listener.handleMessage(serviceBusMessage, broker)
	if err != nil {
		t.Error(err)
	}
}

func TestListener_HandleMessage_InvalidMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	receiver := NewMockReceiver(ctrl)
	listener := NewListener(receiver)
	broker := messaging.NewMockBroker(ctrl)

	serviceBusMessage := &azservicebus.ReceivedMessage{
		Body: []byte("This is a test message"),
	}

	receiver.EXPECT().DeadLetterMessage(context.TODO(), serviceBusMessage, nil).Return(nil)

	err := listener.handleMessage(serviceBusMessage, broker)
	if err != nil {
		t.Error(err)
	}
}
