package service_bus

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"github.com/avast/retry-go"
	"sms-gateway/internal/messaging"
)

type ServiceBusListener struct {
	receiver *azservicebus.Receiver
}

func NewServiceBusListener(receiver *azservicebus.Receiver) *ServiceBusListener {
	return &ServiceBusListener{receiver}
}

func (listener *ServiceBusListener) Listen(callback func(message *messaging.Message)) error {
	for {
		err := retry.Do(
			func() error {
				messages, err := listener.receiver.ReceiveMessages(context.TODO(), 1, nil)
				if err != nil {
					return err
				}

				for _, message := range messages {
					responseMessage, err := toInternalMessage(message)
					if err != nil {
						return err
					}

					callback(responseMessage)

					if err := listener.receiver.CompleteMessage(context.TODO(), message, nil); err != nil {
						return err
					}
				}

				return nil
			}, retry.Attempts(3))
		if err != nil {
			return fmt.Errorf("could not receive message from service bus: %w", err)
		}
	}
}
