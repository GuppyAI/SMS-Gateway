package configuration

import (
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
	"strings"
	"time"
)

var k = koanf.New(".")

func Load() error {
	log.Info().Msg("Loading configuration...")

	defaults := confmap.Provider(map[string]interface{}{
		"logging.level":  "WARN",
		"sms.polling":    5 * time.Second,
		"sms.modem.baud": 115200,
	}, ".")

	if err := k.Load(defaults, nil); err != nil {
		return err
	}

	envProvider := env.Provider("GATEWAY_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "GATEWAY_")), "_", ".", -1)
	})

	if err := k.Load(envProvider, nil); err != nil {
		return err
	}

	allowlist := strings.TrimSpace(k.String("messaging.allowlist"))

	err := k.Set("messaging.allowlist", strings.FieldsFunc(allowlist, func(r rune) bool {
		return r == ','
	}))
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *koanf.Koanf {
	return k
}
