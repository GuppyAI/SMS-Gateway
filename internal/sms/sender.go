package sms

import (
	"github.com/rs/zerolog/log"
	"github.com/warthog618/sms"
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
func (sender *senderImpl) SendSMS(phoneNumber string, message string) error {
	modem := sender.modemProvider.GetModem()

	var logger = log.With().Str("phone_number", phoneNumber).Str("message", message).Logger()

	numberOption := sms.To(phoneNumber)

	pdus, err := sms.Encode([]byte(message), numberOption)
	if err != nil {
		return err
	}

	var binaryPdus [][]byte

	for _, pdu := range pdus {
		binaryPdu, err := pdu.MarshalBinary()
		if err != nil {
			logger.Error().
				Err(err).
				Msg("Error occurred when trying to marshal PDU to binary! Aborting transfer of message...")
			return err
		}

		binaryPdus = append(binaryPdus, binaryPdu)
	}

	logger.Debug().Int("pduCount", len(pdus)).Msg("Proceeding to send message PDUs...")

	for _, binaryPdu := range binaryPdus {
		if _, err := modem.SendPDU(binaryPdu); err != nil {
			logger.Error().Err(err).Msg("Error occurred when trying to send binary pdu! Aborting transfer of remaining PDUs if there are any...")
			return err
		}
	}

	return nil
}
