package main

import (
	"github.com/rs/zerolog/log"
	"sms-gateway/internal/gateway"
)

func main() {

	if err := gateway.Execute(); err != nil {
		log.Warn().Msg("Starting gateway...")
		panic(err)
	}

}
