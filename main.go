package main

import (
	"context"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	lf, err := os.OpenFile(
		"logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("unable to open log file")
	}
	defer lf.Close()

	multiwriters := zerolog.MultiLevelWriter(os.Stdout, lf)
	log.Logger = zerolog.New(multiwriters).With().Timestamp().Logger()

	log := log.With().Str("booking_id", "1234").Logger()
	ctx := log.WithContext(context.Background())
	makeBooking(ctx)

	log.Debug().
		Int("num_ticket", 2).
		Msg("booking is created")
}

func makeBooking(ctx context.Context) {
	log.Ctx(ctx).Info().Msg("creating a booking")
}
