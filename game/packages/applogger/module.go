package applogger

import (
	"go.uber.org/fx"
)

const (
	moduleTag = "mod"
	module    = "logger"
)

var Module = fx.Module(module,
	fx.Provide(
		NewLogger,
	),
)
