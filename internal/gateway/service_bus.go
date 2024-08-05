package gateway

import (
	"fmt"
	"sms-gateway/internal/servicebus"
)

func initializeServiceBus() (servicebus.Sender, servicebus.Receiver, error) {
	provider, err := servicebus.NewProvider()
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
