package logger

import (
	"bytes"
	"log/slog"
	"testing"
)

func TestNewWithoutSentry(t *testing.T) {
	log := New("")
	if log == nil {
		t.Fatal("logger.New(\"\") returned nil")
	}

	// Should not panic when logging
	log.Info("test message")
	log.Warn("test warning")
	log.Error("test error")
}

func TestNewWithInvalidDSN(t *testing.T) {
	// Invalid DSN should panic during Init
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for invalid DSN, but got none")
		}
	}()

	New("invalid-dsn")
}

func TestLoggerOutput(t *testing.T) {
	buf := new(bytes.Buffer)
	handler := slog.NewTextHandler(buf, &slog.HandlerOptions{Level: slog.LevelDebug})
	log := slog.New(handler)

	log.Info("test message", slog.String("key", "value"))

	output := buf.String()
	if output == "" {
		t.Error("expected log output, got empty string")
	}
	if !bytes.Contains(buf.Bytes(), []byte("test message")) {
		t.Errorf("expected 'test message' in output, got: %s", output)
	}
}

func TestLoggerLevels(t *testing.T) {
	tests := []struct {
		name  string
		level slog.Level
	}{
		{"DEBUG", slog.LevelDebug},
		{"INFO", slog.LevelInfo},
		{"WARN", slog.LevelWarn},
		{"ERROR", slog.LevelError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			handler := slog.NewTextHandler(buf, &slog.HandlerOptions{Level: tt.level})
			log := slog.New(handler)

			switch tt.level {
			case slog.LevelDebug:
				log.Debug("debug")
			case slog.LevelInfo:
				log.Info("info")
			case slog.LevelWarn:
				log.Warn("warn")
			case slog.LevelError:
				log.Error("error")
			}

			if buf.Len() == 0 {
				t.Errorf("expected log output for level %v", tt.level)
			}
		})
	}
}
