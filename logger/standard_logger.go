package logger

import (
	"log"
	"os"
)

type StandardLogger struct {
	level       LogLevel
	debugLogger *log.Logger
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

func NewStandardLogger() *StandardLogger {
	return &StandardLogger{
		level:       INFO,
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.LstdFlags),
		infoLogger:  log.New(os.Stdout, "INFO: ", log.LstdFlags),
		warnLogger:  log.New(os.Stdout, "WARN: ", log.LstdFlags),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.LstdFlags),
	}
}

func (l *StandardLogger) Debug(msg string) {
	if l.level <= DEBUG {
		l.debugLogger.Println(msg)
	}
}

func (l *StandardLogger) Info(msg string) {
	if l.level <= INFO {
		l.infoLogger.Println(msg)
	}
}

func (l *StandardLogger) Warn(msg string) {
	if l.level <= WARN {
		l.warnLogger.Println(msg)
	}
}

func (l *StandardLogger) Error(msg string) {
	if l.level <= ERROR {
		l.errorLogger.Println(msg)
	}
}

func (l *StandardLogger) SetLevel(level LogLevel) {
	l.level = level
}
