package logger

import (
	"context"
	"log/slog"
	"os"

	sentry "github.com/getsentry/sentry-go"
	sentryslog "github.com/getsentry/sentry-go/slog"
	slogmulti "github.com/samber/slog-multi"
)

func New(sentryDSN string) *slog.Logger {
	stdHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})

	if sentryDSN == "" {
		return slog.New(stdHandler)
	}

	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	})
	if err != nil {
		panic("sentry.Init failed: " + err.Error())
	}

	sentryHandler := sentryslog.Option{
		EventLevel: []slog.Level{slog.LevelError},
		LogLevel:   []slog.Level{slog.LevelWarn, slog.LevelError},
	}.NewSentryHandler(context.Background())

	return slog.New(slogmulti.Fanout(stdHandler, sentryHandler))
}
