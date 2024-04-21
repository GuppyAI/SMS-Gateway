package configuration

import (
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/v2"
	"github.com/rs/zerolog/log"
	"strings"
)

var k = koanf.New(".")

func Load() error {
	log.Info().Msg("Loading configuration...")

	envProvider := env.Provider("GATEWAY_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(
			strings.TrimPrefix(s, "GATEWAY_")), "_", ".", -1)
	})

	if err := k.Load(envProvider, nil); err != nil {
		return err
	}

	allowlist := k.String("messaging.allowlist")
	err := k.Set("messaging.allowlist", strings.Split(allowlist, ","))
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() *koanf.Koanf {
	return k
}
