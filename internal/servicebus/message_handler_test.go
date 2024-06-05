package servicebus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/messaging"
	"testing"
)

func TestMessageHandler_Handle(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := NewMockSender(ctrl)
	handler, err := NewMessageHandler(sender)
	if err != nil {
		t.Error(err)
		return
	}

	message := messaging.NewMessage(messaging.Request, *messaging.NewAddress(messaging.SMS, "+49123456789"), "This is a test message")

	serviceBusMessage := &azservicebus.Message{
		Body: []byte("This is a test message"),
		ApplicationProperties: map[string]any{
			"address": "sms://+49123456789",
		},
	}

	sender.EXPECT().SendMessage(context.TODO(), serviceBusMessage, nil)

	if err := handler.Handle(*message); err != nil {
		t.Error(err)
		return
	}
}
