package applogger

import (
	"github.com/rs/zerolog"
)

type RunnableLogger struct {
	logger zerolog.Logger
}

func NewRunnableLogger(logger zerolog.Logger) *RunnableLogger {
	return &RunnableLogger{
		logger: logger,
	}
}

func (rl *RunnableLogger) Printf(format string, args ...interface{}) {
	rl.logger.Debug().Str("mod", "runnable").Msgf(format, args...)
}
