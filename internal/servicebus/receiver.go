package servicebus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
)

type Receiver interface {
	ReceiveMessages(ctx context.Context, maxMessages int, options *azservicebus.ReceiveMessagesOptions) ([]*azservicebus.ReceivedMessage, error)
	ReceiveDeferredMessages(ctx context.Context, sequenceNumbers []int64, options *azservicebus.ReceiveDeferredMessagesOptions) ([]*azservicebus.ReceivedMessage, error)
	PeekMessages(ctx context.Context, maxMessageCount int, options *azservicebus.PeekMessagesOptions) ([]*azservicebus.ReceivedMessage, error)
	RenewMessageLock(ctx context.Context, msg *azservicebus.ReceivedMessage, options *azservicebus.RenewMessageLockOptions) error
	Close(ctx context.Context) error
	CompleteMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.CompleteMessageOptions) error
	AbandonMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.AbandonMessageOptions) error
	DeferMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.DeferMessageOptions) error
	DeadLetterMessage(ctx context.Context, message *azservicebus.ReceivedMessage, options *azservicebus.DeadLetterOptions) error
}
