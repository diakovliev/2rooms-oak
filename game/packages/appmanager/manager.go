package appmanager

import (
	"time"

	"github.com/diakovliev/2rooms-oak/packages/applogger"
	"github.com/pior/runnable"
	"github.com/rs/zerolog"
)

var (
	// runnable AppManager timeout
	shutdownTimeout = 20 * time.Second
	// fx timeout
	ShutdownTimeout = shutdownTimeout + 5*time.Second
)

func New(logger zerolog.Logger) runnable.AppManager {
	runnable.SetLogger(applogger.NewRunnableLogger(logger))
	return runnable.NewManager(
		runnable.ManagerShutdownTimeout(shutdownTimeout),
	)
}
