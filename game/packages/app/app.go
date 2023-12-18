package app

import (
	"go.uber.org/fx"
)

var Module = fx.Module("app",
	fx.Provide(
		fx.Annotate(
			newapp,
			fx.OnStart(func(app *impl) {
				app.start()
			}),
			fx.OnStop(func(app *impl) error {
				return app.stop()
			}),
		),
	),
	fx.Populate(new(*impl)),
)
