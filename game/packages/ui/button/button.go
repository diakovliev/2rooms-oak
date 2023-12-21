package button

import (
	"image/color"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/alg/intgeom"
	"github.com/oakmound/oak/v4/collision"
	"github.com/oakmound/oak/v4/entities"
	"github.com/oakmound/oak/v4/event"
	"github.com/oakmound/oak/v4/mouse"
	"github.com/oakmound/oak/v4/render"
	oakscene "github.com/oakmound/oak/v4/scene"
)

type State string

const (
	Disabled State = "disabled"
	Down     State = "down"
	Up       State = "up"
)

func defaultColors() map[State]color.RGBA {
	return map[State]color.RGBA{
		Disabled: {0, 0, 0, 255},
		Up:       {0, 255, 0, 255},
		Down:     {255, 0, 0, 255},
	}
}

func defaultFonts() map[State]*render.Font {
	font := render.DefaultFont()
	return map[State]*render.Font{
		Disabled: font,
		Up:       font,
		Down:     font,
	}
}

func defaultDims() intgeom.Point2 {
	return intgeom.Point2{100, 100}
}

type Button struct {
	*entities.Entity
	dims intgeom.Point2
	// Color in dependance of state
	colors map[State]color.RGBA
	// Font in dependance of state
	fonts map[State]*render.Font
	// Focus flag
	focus bool
	// Text
	text string
	// State
	state State
	// collision label
	label collision.Label
}

func New(ctx *oakscene.Context, text string, opts ...Option) (ret Button) {

	layer := 1000
	drawLayers := []int{layer + 0, layer + 1}

	ret = Button{
		text:   text,
		colors: defaultColors(),
		fonts:  defaultFonts(),
		dims:   defaultDims(),
		focus:  false,
		state:  Up,
		label:  1000,
	}
	for _, opt := range opts {
		opt(&ret)
	}

	sw := render.NewSwitch(string(ret.state), map[string]render.Modifiable{
		string(Up):       render.NewColorBox(ret.dims.X(), ret.dims.Y(), ret.colors[Up]),
		string(Down):     render.NewColorBox(ret.dims.X(), ret.dims.Y(), ret.colors[Down]),
		string(Disabled): render.NewColorBox(ret.dims.X(), ret.dims.Y(), ret.colors[Disabled]),
	})

	ret.Entity = entities.New(
		ctx,
		entities.WithRenderable(sw),
		entities.WithUseMouseTree(true),
		entities.WithDimensions(floatgeom.Point2{float64(ret.dims.X()), float64(ret.dims.Y())}),
		entities.WithLabel(ret.label),
		entities.WithDrawLayers(drawLayers),
		entities.WithChild(
			entities.WithRenderable(render.NewText(ret.text, float64(ret.dims.X()), float64(ret.dims.Y()))),
			entities.WithDrawLayers(drawLayers[1:]),
			entities.WithLabel(ret.label+1),
		),
	)

	event.Bind(ctx, mouse.DragOn, ret.Entity, func(e *entities.Entity, me *mouse.Event) event.Response {
		me.StopPropagation = true
		return event.ResponseNone
	})

	event.Bind(ctx, mouse.PressOn, ret.Entity, func(e *entities.Entity, me *mouse.Event) event.Response {
		me.StopPropagation = true
		sw.Set(string(Down))
		return event.ResponseNone
	})

	event.Bind(ctx, mouse.ReleaseOn, ret.Entity, func(e *entities.Entity, me *mouse.Event) event.Response {
		me.StopPropagation = true
		sw.Set(string(Up))
		return event.ResponseNone
	})

	event.Bind(ctx, mouse.Drag, ret.Entity, func(e *entities.Entity, me *mouse.Event) event.Response {
		me.StopPropagation = true
		sw.Set(string(Up))
		return event.ResponseNone
	})

	return
}
