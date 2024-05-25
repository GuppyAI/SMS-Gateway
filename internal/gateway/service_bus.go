package gateway

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/messaging/azservicebus"
	"sms-gateway/internal/service_bus"
)

func initializeServiceBus() (*azservicebus.Sender, *azservicebus.Receiver, error) {
	provider, err := service_bus.NewProvider()
	if err != nil {
		return nil, nil, fmt.Errorf("initializing service bus provider: %w", err)
	}

	sender, err := provider.GetSender()
	if err != nil {
		return nil, nil, err
	}

	receiver, err := provider.GetReceiver()
	if err != nil {
		return nil, nil, err
	}

	return sender, receiver, nil
}
