package applogger

import (
	"os"

	"github.com/rs/zerolog"
	"go.uber.org/fx/fxevent"
)

type Flags struct {
	UseFxConsoleLogger bool
}

type writer struct {
	logger zerolog.Logger
}

func (w writer) Write(p []byte) (n int, err error) {
	w.logger.Trace().Msg(string(p))
	return len(p), nil
}

func NewFxLogger(logger zerolog.Logger) fxevent.Logger {
	return &fxevent.ConsoleLogger{W: writer{
		logger: logger.With().Str("mod", "fx").Logger(),
	}}
}

func NewAppFxLogger(logger zerolog.Logger, flags Flags) fxevent.Logger {
	if flags.UseFxConsoleLogger {
		return &fxevent.ConsoleLogger{W: os.Stderr}
	}
	return NewFxLogger(logger)
}
