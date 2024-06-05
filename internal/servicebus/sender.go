package servicebus

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"time"
)

type Sender interface {
	NewMessageBatch(ctx context.Context, options *azservicebus.MessageBatchOptions) (*azservicebus.MessageBatch, error)
	SendMessage(ctx context.Context, message *azservicebus.Message, options *azservicebus.SendMessageOptions) error
	SendAMQPAnnotatedMessage(ctx context.Context, message *azservicebus.AMQPAnnotatedMessage, options *azservicebus.SendAMQPAnnotatedMessageOptions) error
	SendMessageBatch(ctx context.Context, batch *azservicebus.MessageBatch, options *azservicebus.SendMessageBatchOptions) error
	ScheduleMessages(ctx context.Context, messages []*azservicebus.Message, scheduledEnqueueTime time.Time, options *azservicebus.ScheduleMessagesOptions) ([]int64, error)
	ScheduleAMQPAnnotatedMessages(ctx context.Context, messages []*azservicebus.AMQPAnnotatedMessage, scheduledEnqueueTime time.Time, options *azservicebus.ScheduleAMQPAnnotatedMessagesOptions) ([]int64, error)
	CancelScheduledMessages(ctx context.Context, sequenceNumbers []int64, options *azservicebus.CancelScheduledMessagesOptions) error
	Close(ctx context.Context) error
}
