package logger

import (
	"github.com/getsentry/sentry-go"
	"time"
)

type SentryLogger struct {
	level LogLevel
}

func NewSentryLogger(dsn string) *SentryLogger {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn,
	}); err != nil {
		panic("sentry.Init: " + err.Error())
	}
	return &SentryLogger{
		level: WARN,
	}
}

func (l *SentryLogger) Debug(msg string) {}
func (l *SentryLogger) Info(msg string)  {}
func (l *SentryLogger) Warn(msg string) {
	if l.level <= WARN {
		sentry.CaptureMessage("WARN: " + msg)
	}
}
func (l *SentryLogger) Error(msg string) {
	if l.level <= ERROR {
		sentry.CaptureMessage("ERROR: " + msg)
	}
}

func (l *SentryLogger) SetLevel(level LogLevel) {
	l.level = level
}

func (l *SentryLogger) Flush() {
	sentry.Flush(2 * time.Second)
}
