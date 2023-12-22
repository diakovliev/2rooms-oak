package main

import (
	"context"
	"os"
	"runtime"

	"github.com/diakovliev/2rooms-oak/packages/scene/debug"
	"github.com/diakovliev/2rooms-oak/packages/scene/mainmenu"
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
	debug.New(logger, win.Window)
	mainmenu.New(logger, win.Window)
	win.Run(ctx)
}
