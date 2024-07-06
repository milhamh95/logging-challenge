package main

import (
	"errors"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("hello world")

	err := errors.New("payment is already expired")
	log.Fatal().
		Err(err).
		Str("payment_id", "123456").
		Str("payment_status", "failed").
		Str("booking_id", "98765").
		Float64("amount", 100.0).
		Msgf("payment failed for booking %s", "234")
}
