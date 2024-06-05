package messaging

import (
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/configuration"
	"sync"
	"testing"
)

func TestBrokerImpl_SendMessage(t *testing.T) {
	controller := gomock.NewController(t)

	handler := NewMockMessageHandler(controller)
	broker := NewBroker(handler)

	channel := NewMockMessageChannel(controller)
	channel.EXPECT().GetSupportedSchema().Return(AddressSchema("test"))

	broker.AddMessageChannel(channel)

	address := NewAddress("test", "testingAddress")
	message := NewMessage(Response, *address, "Message Content")

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()

	channel.EXPECT().SendMessage(*message).Do(func(_ Message) {
		wg.Done()
	})

	broker.Publish(*message)
}

func TestBrokerImpl_ReceiveMessage(t *testing.T) {
	controller := gomock.NewController(t)

	handler := NewMockMessageHandler(controller)
	broker := NewBroker(handler)

	channel := NewMockMessageChannel(controller)
	channel.EXPECT().GetSupportedSchema().Return(AddressSchema("test"))

	broker.AddMessageChannel(channel)

	address := NewAddress("test", "testingAddress")
	message := NewMessage(Request, *address, "Message Content")

	wg := sync.WaitGroup{}
	wg.Add(1)
	defer wg.Wait()

	handler.EXPECT().Handle(*message).Do(func(_ Message) {
		wg.Done()
	})

	broker.Publish(*message)
}

func TestBrokerImpl_ReceiveMessage_AddressNotAllowed(t *testing.T) {
	t.Setenv("GATEWAY_MESSAGING_ALLOWLIST", "test://otherAddress")
	err := configuration.Load()
	if err != nil {
		t.Fatal(err)
	}

	controller := gomock.NewController(t)

	handler := NewMockMessageHandler(controller)
	broker := NewBroker(handler)

	channel := NewMockMessageChannel(controller)
	channel.EXPECT().GetSupportedSchema().Return(AddressSchema("test"))

	broker.AddMessageChannel(channel)

	address := NewAddress("test", "testingAddress")
	message := NewMessage(Request, *address, "Message Content")

	broker.Publish(*message)
}

func TestBrokerImpl_SendMessage_AddressNotAllowed(t *testing.T) {
	t.Setenv("GATEWAY_MESSAGING_ALLOWLIST", "test://otherAddress")
	err := configuration.Load()
	if err != nil {
		t.Fatal(err)
	}

	controller := gomock.NewController(t)

	handler := NewMockMessageHandler(controller)
	broker := NewBroker(handler)

	channel := NewMockMessageChannel(controller)
	channel.EXPECT().GetSupportedSchema().Return(AddressSchema("test"))

	broker.AddMessageChannel(channel)

	address := NewAddress("test", "testingAddress")
	message := NewMessage(Response, *address, "Message Content")

	broker.Publish(*message)
}
