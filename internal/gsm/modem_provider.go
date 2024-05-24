package gsm

import (
	"fmt"
	"github.com/warthog618/modem/at"
)

type ModemProvider interface {
	// GetModem returns the modem in the ModemProvider.
	GetModem() Modem
	// ResetModem can be used to reset the modem into a known state.
	ResetModem() error
}

type modemProviderImpl struct {
	modem Modem
}

func NewModemProvider(modem Modem) (ModemProvider, error) {
	return &modemProviderImpl{modem}, nil
}

func (provider *modemProviderImpl) GetModem() Modem {
	return provider.modem
}

// ResetModem can be used to reset the modem into a known state.
// Will use ATZ to reset to factory defaults and configure the modem to use the SIM SMS storage.
func (provider *modemProviderImpl) ResetModem() error {
	err := provider.modem.Init(at.WithCmds("Z", "+CPMS=SM,SM,SM"))
	if err != nil {
		return fmt.Errorf("error during modem reset: %w", err)
	}

	return nil
}
