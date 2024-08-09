package sms

import (
	"errors"
	"github.com/warthog618/modem/at"
	gsmExt "github.com/warthog618/modem/gsm"
	"go.uber.org/mock/gomock"
	"sms-gateway/internal/configuration"
	"sms-gateway/internal/gsm"
	"testing"
)

func TestReceiverImpl_Listen_Polling(t *testing.T) {
	controller := gomock.NewController(t)
	modemProvider := gsm.NewMockModemProvider(controller)

	receiver, err := NewReceiver(modemProvider)
	if err != nil {
		t.Fatal(err)
	}

	successful := make(chan bool)

	t.Setenv("GATEWAY_SMS_POLLING", "1s")

	if err := configuration.Load(); err != nil {
		t.Fatal(err)
	}

	modem := gsm.NewMockModem(controller)
	modem.EXPECT().AddIndication("+CMT:", gomock.Any(), at.WithTrailingLine)
	modem.EXPECT().Command("+CMGL=4", gomock.Any()).AnyTimes().Return(nil, nil).Do(func(_ string, _ ...at.CommandOption) {
		successful <- true
	})

	modemProvider.EXPECT().GetModem().Return(modem)

	go receiver.Listen(func(message gsmExt.Message) {
		/* DO NOTHING */
	})

	for range successful {
		return
	}
}

func TestReceiverImpl_Listen_PollingError(t *testing.T) {
	controller := gomock.NewController(t)
	modemProvider := gsm.NewMockModemProvider(controller)

	receiver, err := NewReceiver(modemProvider)
	if err != nil {
		t.Fatal(err)
	}

	successful := make(chan bool)

	modem := gsm.NewMockModem(controller)
	modem.EXPECT().AddIndication("+CMT:", gomock.Any(), at.WithTrailingLine)
	modem.EXPECT().Command("+CMGL=4", gomock.Any()).AnyTimes().Return(nil, errors.New("Testing!"))

	modemProvider.EXPECT().ResetModem().AnyTimes().Return(nil)
	modemProvider.EXPECT().GetModem().Return(modem).Do(func() {
		successful <- true
	})

	go receiver.Listen(func(message gsmExt.Message) {
		/* DO NOTHING */
	})

	for range successful {
		return
	}
}

func TestReceiverImpl_Listen_ReceiveMessage(t *testing.T) {
	controller := gomock.NewController(t)
	modemProvider := gsm.NewMockModemProvider(controller)

	receiver, err := NewReceiver(modemProvider)
	if err != nil {
		t.Fatal(err)
	}

	successful := make(chan bool)

	modem := gsm.NewMockModem(controller)
	modem.EXPECT().Command("+CMGL=4", gomock.Any()).AnyTimes()
	modem.EXPECT().AddIndication("+CMT:", gomock.Any(), at.WithTrailingLine).Do(func(_ string, handler func(info []string), _ ...at.IndicationOption) {
		handler([]string{"+CMT: ,54", "0791448720003023240C919471123254760000111011315214002754747A0E4ACF416150BB3C9F87CF6590F92D07D1CB737ADA7D06C1EB72F87B5E9EBB00"})
	})

	modemProvider.EXPECT().GetModem().Return(modem)

	go receiver.Listen(func(message gsmExt.Message) {
		if message.Message != "This is a message for testing purposes." {
			successful <- false
		}

		if message.Number != "+491721234567" {
			successful <- false
		}

		successful <- true
	})

	for success := range successful {
		if success {
			return
		}

		t.FailNow()
	}
}
