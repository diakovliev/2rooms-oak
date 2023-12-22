package button

import (
	"image"
	"image/color"

	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/alg/intgeom"
	"github.com/diakovliev/oak/v4/collision"
	"github.com/diakovliev/oak/v4/entities"
	"github.com/diakovliev/oak/v4/event"
	"github.com/diakovliev/oak/v4/mouse"
	"github.com/diakovliev/oak/v4/render"
	"github.com/diakovliev/oak/v4/render/mod"
	oakscene "github.com/diakovliev/oak/v4/scene"
	"golang.org/x/image/colornames"
)

type State string

const (
	Disabled State = "disabled"
	Down     State = "down"
	Up       State = "up"
)

func defaultColors() map[State]color.RGBA {
	return map[State]color.RGBA{
		Disabled: colornames.Steelblue,
		Up:       colornames.Skyblue,
		Down:     colornames.Blueviolet,
	}
}

func defaultFontColors(colors map[State]color.RGBA) map[State]color.RGBA {
	return map[State]color.RGBA{
		Disabled: utils.InverseColor(colors[Disabled]),
		Up:       utils.InverseColor(colors[Up]),
		Down:     utils.InverseColor(colors[Down]),
	}
}

func defaultFonts(fontColors map[State]color.RGBA) (ret map[State]*render.Font) {
	font := render.DefaultFont()
	fonts := map[State]*render.Font{
		Disabled: font.Copy(),
		Up:       font.Copy(),
		Down:     font.Copy(),
	}
	regenerateFont := func(f *render.Font, clr color.RGBA) (ret *render.Font) {
		ret, _ = f.RegenerateWith(func(fg render.FontGenerator) render.FontGenerator {
			fg.Color = image.NewUniform(clr)
			return fg
		})
		return
	}
	ret = map[State]*render.Font{}
	for k, f := range fonts {
		ret[k] = regenerateFont(f, fontColors[k])
	}
	return
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
	// round x and y
	roundX, roundY float64

	sw      *render.Switch
	swLabel *render.Switch

	callback func()
}

func New(ctx *oakscene.Context, text string, opts ...Option) (ret Button) {

	layer := 1000
	drawLayers := []int{layer + 0, layer + 1}

	colors := defaultColors()
	fontColors := defaultFontColors(colors)

	ret = Button{
		text:   text,
		colors: colors,
		fonts:  defaultFonts(fontColors),
		dims:   defaultDims(),
		focus:  false,
		state:  Up,
		label:  1000,
		roundX: .1,
		roundY: .1,
	}
	for _, opt := range opts {
		opt(&ret)
	}

	ret.sw = render.NewSwitch(string(ret.state), map[string]render.Modifiable{
		string(Up):       render.NewColorBox(ret.dims.X(), ret.dims.Y(), ret.colors[Up]),
		string(Down):     render.NewColorBox(ret.dims.X(), ret.dims.Y(), ret.colors[Down]),
		string(Disabled): render.NewColorBox(ret.dims.X(), ret.dims.Y(), ret.colors[Disabled]),
	})

	entRect := floatgeom.Rect2{
		Min: floatgeom.Point2{0, 0},
		Max: floatgeom.Point2{float64(ret.dims.X()), float64(ret.dims.Y())},
	}

	textArgs := func(text string, font *render.Font) (str string, x float64, y float64) {
		textRect := layout2d.AlignRect(
			layout2d.HCenter|layout2d.VCenter,
			entRect,
			utils.TextMeasureRect(text, font),
			0,
		)
		return text, textRect.Min.X(), textRect.Min.Y()
	}

	ret.swLabel = render.NewSwitch(string(ret.state), map[string]render.Modifiable{
		string(Up):       ret.fonts[Up].NewText(textArgs(ret.text, ret.fonts[Up])).ToSprite(),
		string(Down):     ret.fonts[Down].NewText(textArgs(ret.text, ret.fonts[Down])).ToSprite(),
		string(Disabled): ret.fonts[Disabled].NewText(textArgs(ret.text, ret.fonts[Disabled])).ToSprite(),
	})

	ret.Entity = entities.New(
		ctx,
		entities.WithRenderable(ret.sw),
		entities.WithUseMouseTree(true),
		entities.WithDimensions(floatgeom.Point2{float64(ret.dims.X()), float64(ret.dims.Y())}),
		entities.WithLabel(ret.label),
		entities.WithDrawLayers(drawLayers),
		entities.WithMod(mod.CutRound(ret.roundX, ret.roundY)),
		entities.WithChild(
			entities.WithRenderable(ret.swLabel),
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
		if ret.state == Disabled {
			return event.ResponseNone
		}
		ret.SetState(Down)
		return event.ResponseNone
	})

	event.Bind(ctx, mouse.ReleaseOn, ret.Entity, func(e *entities.Entity, me *mouse.Event) event.Response {
		me.StopPropagation = true
		if ret.state == Disabled {
			return event.ResponseNone
		}
		if ret.callback != nil {
			ret.callback()
		}
		ret.SetState(Up)
		return event.ResponseNone
	})

	event.Bind(ctx, mouse.Drag, ret.Entity, func(e *entities.Entity, me *mouse.Event) event.Response {
		me.StopPropagation = true
		if ret.state == Disabled {
			return event.ResponseNone
		}
		if !e.Rect.Contains(me.Point2) {
			ret.SetState(Up)
		}
		return event.ResponseNone
	})

	return
}

func (b *Button) SetState(state State) {
	b.state = state
	b.swLabel.Set(string(state))
	b.sw.Set(string(state))
}

func (b *Button) Font() *render.Font {
	return b.fonts[b.state]
}
