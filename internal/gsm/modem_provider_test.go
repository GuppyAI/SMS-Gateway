package gsm

import (
	"github.com/stretchr/testify/assert"
	"github.com/warthog618/modem/at"
	"go.uber.org/mock/gomock"
	"testing"
)

func TestNew(t *testing.T) {
	ctrl := gomock.NewController(t)

	gsm := NewMockModem(ctrl)

	smsProvider, err := NewModemProvider(gsm)
	if err != nil {
		t.Error(err)
	}

	assert.IsType(t, &modemProviderImpl{}, smsProvider)
}

func TestProviderImpl_ResetModem(t *testing.T) {
	ctrl := gomock.NewController(t)
	gsm := NewMockModem(ctrl)

	smsProvider, err := NewModemProvider(gsm)
	if err != nil {
		t.Fatal(err)
	}

	gsm.EXPECT().
		Init(at.WithCmds("Z", "+CPMS=SM,SM,SM")).
		Return(nil).
		Times(1)

	err = smsProvider.ResetModem()
	if err != nil {
		t.Fatal(err)
	}
}
