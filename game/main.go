package main

import (
	"context"
	"os"
	"runtime"

	"github.com/diakovliev/2rooms-oak/packages/scene/debug"
	"github.com/diakovliev/2rooms-oak/packages/scene/initial"
	"github.com/diakovliev/2rooms-oak/packages/window"
	"github.com/rs/zerolog"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	logger := zerolog.New(os.Stdout)
	ctx := context.Background()
	win := window.New()
	debug.New(logger, win)
	initial.New(logger, win)
	win.Run(ctx)
}
