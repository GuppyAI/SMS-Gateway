package gateway

import (
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
	"github.com/warthog618/modem/at"
	gsmExt "github.com/warthog618/modem/gsm"
	modemSerial "github.com/warthog618/modem/serial"
	"github.com/warthog618/modem/trace"
	"io"
	"sms-gateway/internal/application_context"
	"sms-gateway/internal/configuration"
	"sms-gateway/internal/echo"
	"sms-gateway/internal/gsm"
	"sms-gateway/internal/logging"
	"sms-gateway/internal/messaging"
	"sms-gateway/internal/sms"
)

var config *koanf.Koanf

// Execute will start the application
func Execute() error {
	if err := configuration.Load(); err != nil {
		log.Err(err)
		return err
	}

	if err := logging.Setup(); err != nil {
		log.Err(err)
		return err
	}

	config = configuration.GetConfig()

	log.Debug().Any("config", config.All()).Msg("Configuration")

	if len(config.Strings("messaging.allowlist")) == 0 {
		log.Warn().Msg("STARTING WITHOUT AN ALLOWLIST! THIS IS PROBABLY NOT INTENDED!")
	}

	broker := messaging.NewBroker(echo.New())

	application_context.Init(broker)

	if err := initializeSMS(broker); err != nil {
		log.Error().Err(err).Msg("Error occurred during initialization of SMS channel")
		return err
	}

	return nil
}

// initializeSMS initializes the SMS channel and registers it with the messaging broker
func initializeSMS(broker messaging.Broker) error {
	device := config.String("sms.modem.device")
	baud := config.Int("sms.modem.baud")

	serialPort, err := modemSerial.New(modemSerial.WithPort(device), modemSerial.WithBaud(baud))
	if err != nil {
		log.Error().
			Err(err).
			Str("device", device).
			Int("baud", baud).
			Msgf("Error occurred while trying to access device!")

		return err
	}

	tracing := config.Bool("sms.tracing")

	var modemIO io.ReadWriter = serialPort

	if tracing {
		modemIO = trace.New(serialPort, trace.WithLogger(logging.NewModemTracer(log.With().Logger())))
	}

	modem := gsmExt.New(at.New(modemIO))

	modemProvider, err := gsm.NewModemProvider(modem)
	if err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to initialize sms provider")
		return err
	}

	if err := modemProvider.ResetModem(); err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to reset modem")
		return err
	}

	sender, err := sms.NewSender(modemProvider)
	if err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to create SMS sender")
		return err
	}

	receiver, err := sms.NewReceiver(modemProvider)
	if err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to create SMS receiver")
		return err
	}

	messageChannel := sms.NewMessageChannel(sender, receiver)

	broker.AddMessageChannel(messageChannel)

	return nil
}
