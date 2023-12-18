package window

import (
	"context"

	"github.com/diakovliev/2rooms-oak/packages/scene"
	"github.com/oakmound/oak/v4"
	"github.com/pior/runnable"
)

type Window struct {
	*oak.Window
	initialScene string
}

func New(appManager runnable.AppManager) (ret *Window) {
	ret = &Window{
		Window:       oak.NewWindow(),
		initialScene: scene.Initial,
	}
	appManager.Add(ret)
	return
}

func (w Window) Run(ctx context.Context) error {
	w.Window.Init(w.initialScene)
	return context.Canceled
}
