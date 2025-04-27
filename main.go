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
	r.HandleFunc("/", handler)

	// start: set up any of your logger configuration here if necessary
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	logFile, err := os.OpenFile(
		"logs/app.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0666,
	)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to open log file")
	}

	multiWriters := zerolog.MultiLevelWriter(os.Stdout, logFile)
	log.Logger = zerolog.New(multiWriters).With().Timestamp().Logger()
	log.Info().Msg("hello world")

	// end: set up any of your logger configuration here

	server := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("failed to listen and serve http server")
		}
	}()
	<-ctx.Done()

	if err := server.Shutdown(context.Background()); err != nil {
		log.Error().Err(err).Msg("failed to shutdown http server gracefully")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	clog := log.With().
		Str("request_id", uuid.New().String()).
		Str("function_name", "handler").
		Logger()

	ctx := clog.WithContext(r.Context())
	name := r.URL.Query().Get("name")
	res, err := greeting(ctx, name)
	if err != nil {
		clog.Error().Ctx(ctx).Err(err).Msg("request failed")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	clog.Info().Ctx(ctx).Msg("request received")
	w.Write([]byte(res))
}

func greeting(ctx context.Context, name string) (string, error) {
	log.Ctx(ctx).Debug().Str("function_name", "greeting").Msg("do greeting")
	if len(name) < 5 {
		return fmt.Sprintf("Hello %s! Your name is to short\n", name), nil
	}
	return fmt.Sprintf("Hi %s", name), nil
}
