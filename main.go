package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	log.Info().Msg("hello world")
	log.Debug().Msg("asdf")
	log.Warn().Msg("asdf")
	log.Error().Msg("asdf")
	log.Fatal().Msg("asdf")
}
