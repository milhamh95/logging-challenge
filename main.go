package main

import (
	"context"

	"github.com/rs/zerolog/log"
)

func main() {
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
