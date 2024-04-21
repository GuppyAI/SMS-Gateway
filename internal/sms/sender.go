package sms

import (
	"github.com/rs/zerolog/log"
	"sms-gateway/internal/gsm"
)

// A Sender can be used to send SMS messages.
type Sender interface {
	// SendSMS sends an SMS containing the given message to the given phone number.
	SendSMS(phoneNumber string, message string) error
}

// senderImpl is the default implementation of a Sender.
type senderImpl struct {
	// modemProvider provides the Modem to send SMS.
	modemProvider gsm.ModemProvider
}

// NewSender can be used to construct a new Sender from a ModemProvider
func NewSender(modemProvider gsm.ModemProvider) (Sender, error) {
	return &senderImpl{modemProvider: modemProvider}, nil
}

// SendSMS uses the given Modem to send an SMS.
// Will use the Modem.SendShortMessage method when the message has a maximum of 160 characters and the Modem.SendLongMessage otherwise.
func (sender *senderImpl) SendSMS(phoneNumber string, message string) error {
	modem := sender.modemProvider.GetModem()

	var logger = log.With().Str("phone_number", phoneNumber).Str("message", message).Logger()

	var err error

	if len(message) <= 160 {
		logger.Debug().Msg("Sending short message")
		_, err = modem.SendShortMessage(phoneNumber, message)
	} else {
		logger.Debug().Msg("Sending long message")
		_, err = modem.SendLongMessage(phoneNumber, message)
	}

	if err != nil {
		logger.Error().
			Err(err).
			Msgf("Error occurred while trying to send SMS message")

		return err
	}

	return nil
}
