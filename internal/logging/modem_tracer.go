package logging

import (
	"fmt"
	"github.com/rs/zerolog"
	"strings"
)

const (
	WRITE_PREFIX = "w:"
	READ_PREFIX  = "r:"
)

type ModemTracerWriter struct {
	logger zerolog.Logger
}

func NewModemTracer(logger zerolog.Logger) *ModemTracerWriter {
	return &ModemTracerWriter{logger}
}

func (m ModemTracerWriter) Printf(format string, v ...interface{}) {
	formatted := fmt.Sprintf(format, v...)

	replace := []string{" ", "\r", "\n"}

	for _, r := range replace {
		formatted = strings.ReplaceAll(formatted, r, "")
	}

	if strings.HasPrefix(formatted, WRITE_PREFIX) {
		m.logger.Info().Str("written", strings.TrimPrefix(formatted, WRITE_PREFIX)).Msg("Modem command was sent")
		return
	}

	if strings.HasPrefix(formatted, READ_PREFIX) {
		m.logger.Info().Str("read", strings.TrimPrefix(formatted, READ_PREFIX)).Msg("Modem has sent answer")
		return
	}

	m.logger.Warn().Str("trace", formatted).Msg("Got trace with unknown signature")
}
