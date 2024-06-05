package gateway

import (
	"github.com/rs/zerolog/log"
	"github.com/warthog618/modem/at"
	gsmExt "github.com/warthog618/modem/gsm"
	modemSerial "github.com/warthog618/modem/serial"
	"github.com/warthog618/modem/trace"
	"io"
	"sms-gateway/internal/gsm"
	"sms-gateway/internal/logging"
	"sms-gateway/internal/messaging"
	"sms-gateway/internal/sms"
)

// initializeSMSChannel initializes the SMS channel
func initializeSMSChannel() (messaging.MessageChannel, error) {
	device := config.String("sms.modem.device")
	baud := config.Int("sms.modem.baud")

	serialPort, err := modemSerial.New(modemSerial.WithPort(device), modemSerial.WithBaud(baud))
	if err != nil {
		log.Error().
			Err(err).
			Str("device", device).
			Int("baud", baud).
			Msgf("Error occurred while trying to access device!")

		return nil, err
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
		return nil, err
	}

	if err := modemProvider.ResetModem(); err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to reset modem")
		return nil, err
	}

	sender, err := sms.NewSender(modemProvider)
	if err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to create SMS sender")
		return nil, err
	}

	receiver, err := sms.NewReceiver(modemProvider)
	if err != nil {
		log.Error().Err(err).Msg("Error occurred while trying to create SMS receiver")
		return nil, err
	}

	return sms.NewMessageChannel(sender, receiver), nil
}
