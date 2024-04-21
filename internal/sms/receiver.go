package sms

import (
	"github.com/rs/zerolog/log"
	"github.com/warthog618/modem/at"
	gsmExt "github.com/warthog618/modem/gsm"
	"github.com/warthog618/sms"
	"github.com/warthog618/sms/encoding/tpdu"
	"sms-gateway/internal/configuration"
	"sms-gateway/internal/gsm"
	"time"
)

// A Receiver can be used to receive SMS messages.
type Receiver interface {
	// Listen listens for SMS messages and hands them over to the given handler.
	Listen(handler func(gsmExt.Message))
}

// receiverImpl is the default implementation of a Receiver.
type receiverImpl struct {
	modemProvider gsm.ModemProvider
	collector     *sms.Collector
}

// NewReceiver constructs a new Receiver from a Modem
func NewReceiver(modemProvider gsm.ModemProvider) (Receiver, error) {
	return &receiverImpl{
		modemProvider: modemProvider,
		collector: sms.NewCollector(sms.WithReassemblyTimeout(time.Minute, func(tpdus []*tpdu.TPDU) {
			log.Error().
				Err(gsmExt.ErrReassemblyTimeout{TPDUs: tpdus}).
				Msg("Reassembly of messages timed out!")
		})),
	}, nil
}

// Listen polls the Modem for new messages.
// The polling frequency can be configured using the sms.polling_timeout configuration option.
func (receiver *receiverImpl) Listen(handler func(gsmExt.Message)) {
	modem := receiver.modemProvider.GetModem()

	err := modem.AddIndication("+CMT:", func(info []string) {
		log.Debug().Msg("Receiving message!")

		pdu, err := gsmExt.UnmarshalTPDU(info)
		if err != nil {
			log.Error().
				Err(err).
				Strs("message_info", info).
				Msg("Error occurred on reception of SMS message")

			return
		}

		tpdus, err := receiver.collector.Collect(pdu)
		if err != nil {
			log.Error().
				Err(err).
				Any("pdu", pdu).
				Msg("Error occurred while trying to collect TPDUs")

			return
		}

		if tpdus == nil {
			log.Debug().Msg("Skipping message handling, concatenated SMS is still collecting...")
			return
		}

		message, err := sms.Decode(tpdus)
		if err != nil {
			log.Error().
				Err(err).
				Any("tpdus", tpdus).
				Msg("Error occurred while trying to decode TPDUs")

			return
		}

		gsmMessage := gsmExt.Message{
			Number:  tpdus[0].OA.Number(),
			Message: string(message),
			SCTS:    tpdus[0].SCTS,
			TPDUs:   tpdus,
		}

		log.Debug().
			Str("number", gsmMessage.Number).
			Str("message", gsmMessage.Message).
			Msg("Successfully parsed SMS message!")

		handler(gsmMessage)
	}, at.WithTrailingLine)

	if err != nil {
		log.Error().
			Err(err).
			Msg("Error occurred while trying to add indication to modem")

		return
	}

	config := configuration.GetConfig()
	pollingTimeout := config.Duration("sms.polling_timeout") * time.Millisecond

	for {
		select {
		case <-time.After(pollingTimeout):
			_, err := modem.Command("+CMGL=4")
			if err != nil {
				log.Error().
					Err(err).
					Msg("Error occurred while polling for new messages. Resetting modem...")

				resetErr := receiver.modemProvider.ResetModem()
				if resetErr != nil {
					log.Error().
						Err(resetErr).
						Msg("Could not reset modem after error occurred while polling for new messages.")
				}
			}
		}
	}
}
