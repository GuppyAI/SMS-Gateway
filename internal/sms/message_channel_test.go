package sms

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/messaging"
	"testing"
)

func TestNewMessageChannel(t *testing.T) {
	ctrl := gomock.NewController(t)

	sender := NewMockSender(ctrl)
	receiver := NewMockReceiver(ctrl)

	messageChannel := NewMessageChannel(sender, receiver)

	assert.IsType(t, &MessageChannel{}, messageChannel)
}

func TestMessageChannel_GetSupportedSchema(t *testing.T) {
	ctrl := gomock.NewController(t)

	sender := NewMockSender(ctrl)
	receiver := NewMockReceiver(ctrl)

	messageChannel := NewMessageChannel(sender, receiver)
	supportedSchema := messageChannel.GetSupportedSchema()

	assert.Equal(t, messaging.SMS, supportedSchema)
}

func TestMessageChannel_ReceiveMessages(t *testing.T) {
	ctrl := gomock.NewController(t)

	sender := NewMockSender(ctrl)
	receiver := NewMockReceiver(ctrl)
	broker := messaging.NewMockBroker(ctrl)

	messageChannel := NewMessageChannel(sender, receiver)

	receiver.EXPECT().Listen(gomock.Any()).Times(1)

	err := messageChannel.ReceiveMessages(broker)
	if err != nil {
		t.Fatal(err)
	}
}
