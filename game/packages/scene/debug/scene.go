package debug

import (
	"runtime"

	"github.com/diakovliev/2rooms-oak/packages/scene/scenes"
	"github.com/diakovliev/2rooms-oak/packages/ui/debuginfo"
	"github.com/oakmound/oak/v4"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/key"
	"github.com/oakmound/oak/v4/mouse"
	"github.com/oakmound/oak/v4/render"
	oakscene "github.com/oakmound/oak/v4/scene"
	"github.com/rs/zerolog"
)

type Scene struct {
	logger zerolog.Logger
	Name   string
	End    string
}

func New(logger zerolog.Logger, w *oak.Window) (ret *Scene) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	ret = &Scene{
		logger: logger,
		Name:   scenes.Debug,
		End:    scenes.MainMenu,
	}
	if err := w.AddScene(ret.Name, ret.build()); err != nil {
		logger.Panic().Err(err).Msgf("failed to add '%s' scene", ret.Name)
	}
	return
}

func (s Scene) start(ctx *oakscene.Context) {
	debuginfo.DebugInfo(ctx)
	ctx.DrawStack.Draw(render.NewText("Debug scene! Any key to return to in initial", 100, 100))
	event.GlobalBind(ctx, key.AnyDown, func(key.Event) event.Response {
		ctx.Window.GoToScene(s.End)
		return 0
	})
	event.GlobalBind(ctx, mouse.Press, func(me *mouse.Event) event.Response {
		ctx.Window.GoToScene(s.End)
		return event.ResponseNone
	})
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
