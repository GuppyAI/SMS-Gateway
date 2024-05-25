package gateway

import (
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
	"slices"
	"sms-gateway/internal/configuration"
	"sms-gateway/internal/logging"
	"sms-gateway/internal/messaging"
	"sms-gateway/internal/service_bus"
	"strings"
)

var config *koanf.Koanf

// Execute will start the application
func Execute() error {
	if err := configuration.Load(); err != nil {
		log.Err(err).Msg("Could not load configuration!")
		return err
	}

	if err := logging.Setup(); err != nil {
		log.Err(err).Msg("Could not setup logging!")
		return err
	}

	config = configuration.GetConfig()
	if err := configuration.ValidateConfig(); err != nil {
		log.Err(err).Msg("Invalid configuration!")
		return err
	}

	log.Debug().Any("config", config.All()).Msg("Starting up...")

	sender, receiver, err := initializeServiceBus()
	if err != nil {
		return err
	}

	handler, err := service_bus.NewMessageHandler(sender)
	if err != nil {
		return err
	}

	broker := messaging.NewBroker(handler)

	errorChannel := make(chan error)

	if err := setupMessageChannels(errorChannel, broker); err != nil {
		log.Error().Err(err).Msg("Could not setup message channels")
		return err
	}

	if err := service_bus.NewServiceBusListener(receiver).Listen(func(message *messaging.Message) {
		broker.Publish(*message)
	}); err != nil {
		return err
	}

	for {
		select {
		case err := <-errorChannel:
			log.Error().Err(err).Msg("Error occurred in message channel!")
			return err
		}
	}
}

func setupMessageChannels(errorChannel chan error, broker messaging.Broker) error {
	messageChannelTypes := strings.Split(config.String("messaging.channels"), ",")
	slices.Sort(messageChannelTypes)
	messageChannelTypes = slices.Compact(messageChannelTypes)

	for _, messageChannelType := range messageChannelTypes {
		var messageChannel messaging.MessageChannel

		switch strings.ToLower(messageChannelType) {
		case "sms":
			var err error
			messageChannel, err = initializeSMSChannel()
			if err != nil {
				log.Error().Err(err).Msg("Could not instantiate SMS channel!")
				return err
			}
		}

		broker.AddMessageChannel(messageChannel)
		go func() {
			if err := messageChannel.ReceiveMessages(broker); err != nil {
				errorChannel <- err
			}
		}()
	}

	return nil
}
