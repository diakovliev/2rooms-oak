package debug

import "go.uber.org/fx"

var Module = fx.Module("debug",
	fx.Provide(New),
	fx.Populate(new(*Scene)),
)
