package window

import (
	"context"
	"runtime"

	"github.com/diakovliev/2rooms-oak/packages/scene"
	"github.com/oakmound/oak/v4"
)

type Window struct {
	*oak.Window
	initialScene string
}

func New() (ret *Window) {
	ret = &Window{
		Window:       oak.NewWindow(),
		initialScene: scene.Initial,
	}
	return
}

func (w Window) Run(ctx context.Context) error {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	w.Window.Init(
		w.initialScene,
		func(c oak.Config) (oak.Config, error) {
			c.Title = "2rooms-oak"
			c.Screen.Width = 640
			c.Screen.Height = 480
			return c, nil
		},
	)
	return context.Canceled
}
