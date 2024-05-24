package sms

import (
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/gsm"
	"testing"
)

func TestSenderImpl_SendSMS_Short(t *testing.T) {
	controller := gomock.NewController(t)
	modemProvider := gsm.NewMockModemProvider(controller)

	modem := gsm.NewMockModem(controller)
	modemProvider.EXPECT().GetModem().Return(modem)

	sender, err := NewSender(modemProvider)
	if err != nil {
		t.Fatal(err)
	}

	message := "This is a short test message"
	phoneNumber := "+493023125000"

	modem.EXPECT().SendPDU(gomock.Any(), gomock.Any()).Times(1)

	if err := sender.SendSMS(phoneNumber, message); err != nil {
		t.Fatal(err)
	}
}

func TestSenderImpl_SendSMS_Long(t *testing.T) {
	controller := gomock.NewController(t)
	modemProvider := gsm.NewMockModemProvider(controller)

	modem := gsm.NewMockModem(controller)
	modemProvider.EXPECT().GetModem().Return(modem)

	sender, err := NewSender(modemProvider)
	if err != nil {
		t.Fatal(err)
	}

	message := "This is a long test message! Some random characters follow: abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	phoneNumber := "+493023125000"

	modem.EXPECT().SendPDU(gomock.Any(), gomock.Any()).Times(2)

	if err := sender.SendSMS(phoneNumber, message); err != nil {
		t.Fatal(err)
	}
}
