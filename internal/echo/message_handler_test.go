package echo

import (
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/application_context"
	"sms-gateway/internal/messaging"
	"testing"
)

func TestMessageHandler_Handle(t *testing.T) {
	controller := gomock.NewController(t)
	broker := messaging.NewMockBroker(controller)

	application_context.Init(broker)

	handler := New()

	address := messaging.NewAddress("test", "testingAddress")
	message := messaging.NewMessage(messaging.Request, *address, "Message Content")

	expectedMessage := messaging.NewMessage(messaging.Response, *address, "Message Content")

	broker.EXPECT().Publish(*expectedMessage).Times(1)

	if err := handler.Handle(*message); err != nil {
		t.Fatal(err)
	}
}
