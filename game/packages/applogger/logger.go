package applogger

import (
	"fmt"
	"os"
	"path"
	"time"

	"github.com/rs/zerolog"
)

const (
	defaultLogFile            = "stdout"
	defaultLogLevel           = "trace"
	defaultLogTimestampFormat = time.RFC3339Nano
)

func NewLogger() zerolog.Logger {
	zerolog.CallerMarshalFunc = func(pc uintptr, file string, line int) string {
		return fmt.Sprintf("%s:%d", path.Base(file), line)
	}
	return zerolog.New(os.Stdout)
}
