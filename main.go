package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, syscall.SIGTERM)
	go func() {
		oscall := <-ch
		log.Warn().Msgf("system call:%+v", oscall)
		cancel()
	}()

	r := mux.NewRouter()
	r.Use(middleware)
	r.HandleFunc("/", handler)

	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	// start: set up any of your logger configuration here if necessary
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

	// end: set up any of your logger configuration here

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		log.Info().Msg("starting the server")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve http server")
		}
	}()
	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server gracefully")
	}
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := log.Logger.With().
			Str("request_id", uuid.New().String()).
			Str("path", r.URL.String()).
			Str("method", r.Method).
			Logger()
		ctx := log.WithContext(r.Context())
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log := log.Ctx(ctx).With().Str("func", "handler").Logger()

	name := r.URL.Query().Get("name")
	log.Debug().
		Str("name", name).
		Msg("processing request to handler endpoint")

	res, err := greeting(ctx, name)
	if err != nil {
		log.Error().Err(err).Msg("greeting failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	log.Info().Msg("processing is finished")
	w.Write([]byte(res))
}

func greeting(ctx context.Context, name string) (string, error) {
	log := log.Ctx(ctx).With().Str("func", "greeting").Logger()
	if len(name) < 5 {
		log.Warn().Msg("name is too short")
		return fmt.Sprintf("Hello %s! Your name is to short\n", name), nil
	}

	log.Info().Msg("name is accepted")
	return fmt.Sprintf("Hi %s", name), nil
}
