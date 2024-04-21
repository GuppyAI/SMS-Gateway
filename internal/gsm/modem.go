package gsm

import "github.com/warthog618/modem/at"

// Modem is an interface which is implicitly implemented by warthog618's gsm.GSM.
// It is primarily used to simplify testing
type Modem interface {
	AddIndication(prefix string, handler at.InfoHandler, options ...at.IndicationOption) error
	CancelIndication(prefix string)
	Closed() <-chan struct{}
	Command(cmd string, options ...at.CommandOption) ([]string, error)
	Escape(b ...byte)
	Init(options ...at.InitOption) error
	SendShortMessage(number string, message string, options ...at.CommandOption) (string, error)
	SendLongMessage(number string, message string, options ...at.CommandOption) ([]string, error)
	SendPDU(tpdu []byte, options ...at.CommandOption) (string, error)
}
