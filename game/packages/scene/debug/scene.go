package debug

import (
	"github.com/diakovliev/2rooms-oak/packages/scene"
	"github.com/diakovliev/2rooms-oak/packages/window"
	"github.com/oakmound/oak/v4/render"
	oakscene "github.com/oakmound/oak/v4/scene"
	"github.com/rs/zerolog"
)

type Scene struct {
	logger zerolog.Logger
	Name   string
	End    string
}

func New(logger zerolog.Logger, w *window.Window) (ret *Scene) {
	ret = &Scene{
		logger: logger,
		Name:   scene.Debug,
		End:    scene.Initial,
	}
	if err := w.AddScene(ret.Name, ret.build()); err != nil {
		logger.Panic().Err(err).Msgf("failed to add '%s' scene", ret.Name)
	}
	return
}

func (s Scene) start(ctx *oakscene.Context) {
	ctx.DrawStack.Draw(render.NewText("Debug scene! Any key to return to in initial", 100, 100))
	// event.GlobalBind(ctx, key.AnyDown, func(key.Event) event.Response {
	// 	ctx.Window.GoToScene(s.End)
	// 	return 0
	// })
}

func (s Scene) end() (string, *oakscene.Result) {
	return s.End, nil
}

func (s Scene) build() oakscene.Scene {
	return oakscene.Scene{
		Start: s.start,
		End:   s.end,
	}
}
