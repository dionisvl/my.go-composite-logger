package main

import (
	"project-root/internal/config"
	"project-root/internal/logger"

	sentryhttp "github.com/getsentry/sentry-go/http"
	"log"
	"net/http"
	"time"
)

func main() {
	config.LoadEnv()

	stdLogger := logger.NewStandardLogger()
	var loggers []logger.Logger
	loggers = append(loggers, stdLogger)

	var sentryHandler *sentryhttp.Handler
	if config.IsSentryEnabled() {
		sentryLogger := logger.NewSentryLogger(config.GetSentryDSN())
		defer sentryLogger.Flush()
		loggers = append(loggers, sentryLogger)

		sentryHandler = sentryhttp.New(sentryhttp.Options{
			Repanic:         true,
			WaitForDelivery: true,
			Timeout:         2 * time.Second,
		})
		log.Println("Sentry enabled and logger initialized successfully.")
	}

	compositeLogger := logger.NewCompositeLogger(loggers)

	compositeLogger.SetLevel(logger.DEBUG)
	compositeLogger.Info("App started")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		compositeLogger.Info("1_helloInfo")
		compositeLogger.Warn("2_hellowarning")
		w.Write([]byte("Hello, World!"))
	})

	if sentryHandler != nil {
		http.ListenAndServe(":8080", sentryHandler.Handle(mux))
	} else {
		http.ListenAndServe(":8080", mux)
	}
}
