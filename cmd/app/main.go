package main

import (
	"log/slog"
	"net/http"
	"project-root/internal/config"
	"project-root/internal/logger"
	"project-root/internal/security"
	"time"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
)

func main() {
	config.LoadEnv()

	log := logger.New(config.GetSentryDSN())
	defer sentry.Flush(2 * time.Second)

	log.Info("App started")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Info("1_helloInfo")
		log.Warn("2_hellowarning")
		if _, err := w.Write([]byte("Hello, World!")); err != nil {
			log.Error("failed to write response", slog.Any("error", err))
		}
	})

	// Stack middleware: rate limit -> body size -> logging
	var handler http.Handler = mux
	handler = security.RateLimitMiddleware(10)(handler)
	handler = security.MaxBodySizeMiddleware(security.MaxBodySize)(handler)
	handler = security.LoggingMiddleware(log)(handler)

	// Wrap with Sentry HTTP handler if enabled
	var finalHandler http.Handler = handler
	if config.IsSentryEnabled() {
		sentryHandler := sentryhttp.New(sentryhttp.Options{
			Repanic:         true,
			WaitForDelivery: true,
			Timeout:         2 * time.Second,
		})
		finalHandler = sentryHandler.Handle(handler)
	}

	server := &http.Server{
		Addr:         ":8080",
		Handler:      finalHandler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Error("server error", slog.Any("error", err))
	}
}
