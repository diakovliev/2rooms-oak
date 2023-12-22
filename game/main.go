package main

import (
	"context"
	"os"
	"runtime"

	"github.com/diakovliev/2rooms-oak/packages/scene/mainmenu"
	"github.com/diakovliev/2rooms-oak/packages/scene/tests/galign"
	"github.com/diakovliev/2rooms-oak/packages/scene/tests/halign"
	"github.com/diakovliev/2rooms-oak/packages/scene/tests/valign"
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
	mainmenu.New(logger, win.Window)

	// Tests
	valign.New(logger, win.Window)
	halign.New(logger, win.Window)
	galign.New(logger, win.Window)

	win.Run(ctx)
}
