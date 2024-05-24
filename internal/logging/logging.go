package logging

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"sms-gateway/internal/configuration"
	"strings"
)

var k = configuration.GetConfig()

func Setup() error {
	log.Info().Msg("Setting up logging...")

	log.Logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Caller()
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level := k.String("logging.level")

	switch strings.ToUpper(level) {
	case "TRACE":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "DEBUG":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "INFO":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "WARN":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "ERROR":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "FATAL":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "NONE":
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	}

	return nil
}
