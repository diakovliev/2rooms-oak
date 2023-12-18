package window

import "go.uber.org/fx"

var Module = fx.Module("window",
	fx.Provide(New),
)
