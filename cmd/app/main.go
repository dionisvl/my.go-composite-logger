package main

import (
	"net/http"
	"project-root/internal/config"
	"project-root/internal/logger"
	"project-root/internal/security"
	"time"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/getsentry/sentry-go"
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
		w.Write([]byte("Hello, World!"))
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

	http.ListenAndServe(":8080", finalHandler)
}
