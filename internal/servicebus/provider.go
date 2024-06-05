package servicebus

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"sms-gateway/internal/configuration"
)

type Provider interface {
	GetReceiver() (Receiver, error)
	GetSender() (Sender, error)
}

type providerImpl struct {
	receiverClient *azservicebus.Client
	senderClient   *azservicebus.Client
}

func NewProvider() (Provider, error) {
	config := configuration.GetConfig()

	senderConnectionString := config.String("servicebus.sender.connectionstring")
	senderClient, err := azservicebus.NewClientFromConnectionString(senderConnectionString, nil)
	if err != nil {
		return nil, fmt.Errorf("creating service bus sender client: %w", err)
	}

	receiverConnectionString := config.String("servicebus.receiver.connectionstring")
	receiverClient, err := azservicebus.NewClientFromConnectionString(receiverConnectionString, nil)
	if err != nil {
		return nil, fmt.Errorf("creating service bus receiver client: %w", err)
	}

	return &providerImpl{
		senderClient:   senderClient,
		receiverClient: receiverClient,
	}, nil
}

func (p *providerImpl) GetReceiver() (Receiver, error) {
	config := configuration.GetConfig()
	queue := config.String("servicebus.receiver.queue")

	return p.receiverClient.NewReceiverForQueue(queue, nil)
}

func (p *providerImpl) GetSender() (Sender, error) {
	config := configuration.GetConfig()
	queue := config.String("servicebus.sender.queue")

	return p.senderClient.NewSender(queue, nil)
}
