package logger

type CompositeLogger struct {
	loggers []Logger
}

func NewCompositeLogger(loggers []Logger) *CompositeLogger {
	return &CompositeLogger{
		loggers: loggers,
	}
}

func (c *CompositeLogger) Debug(msg string) {
	for _, logger := range c.loggers {
		logger.Debug(msg)
	}
}

func (c *CompositeLogger) Info(msg string) {
	for _, logger := range c.loggers {
		logger.Info(msg)
	}
}

func (c *CompositeLogger) Warn(msg string) {
	for _, logger := range c.loggers {
		logger.Warn(msg)
	}
}

func (c *CompositeLogger) Error(msg string) {
	for _, logger := range c.loggers {
		logger.Error(msg)
	}
}

func (c *CompositeLogger) SetLevel(level LogLevel) {
	for _, logger := range c.loggers {
		logger.SetLevel(level)
	}
}
