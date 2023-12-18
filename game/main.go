package main

import (
	"context"

	"github.com/diakovliev/2rooms-oak/packages/app"
	"github.com/diakovliev/2rooms-oak/packages/applogger"
	"github.com/diakovliev/2rooms-oak/packages/appmanager"
	"github.com/diakovliev/2rooms-oak/packages/scene/debug"
	"github.com/diakovliev/2rooms-oak/packages/scene/initial"
	"github.com/diakovliev/2rooms-oak/packages/window"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.WithLogger(applogger.NewAppFxLogger),
		fx.Provide(
			appmanager.New,
			context.Background,
			func() applogger.Flags {
				return applogger.Flags{
					UseFxConsoleLogger: false,
				}
			}),
		applogger.Module,
		debug.Module,
		initial.Module,
		window.Module,
		app.Module,
	).Run()
}
