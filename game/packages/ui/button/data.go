package button

import (
	"image/color"

	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/alg/intgeom"
	"github.com/diakovliev/oak/v4/render"
)

type data struct {
	// button size
	size intgeom.Point2
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
	// round x and y
	roundX, roundY float64

	sw      *render.Switch
	swLabel *render.Switch

	callback func()
}

func newData(text string) *data {
	colors := defaultColors()
	return &data{
		text:   text,
		colors: colors,
		fonts:  defaultFonts(defaultFontColors(colors)),
		size:   defaultSize(),
		focus:  false,
		state:  Up,
		roundX: .1,
		roundY: .1,
	}
}

func (bd data) dimensions() floatgeom.Point2 {
	return floatgeom.Point2{float64(bd.size.X()), float64(bd.size.Y())}
}

func (bd data) bounds() floatgeom.Rect2 {
	return floatgeom.Rect2{
		Min: floatgeom.Point2{0, 0},
		Max: bd.dimensions(),
	}
}

func (bd data) isState(state State) bool {
	return bd.state == state
}

func (bd *data) setState(state State) {
	bd.state = state
	bd.sw.Set(string(state))
	bd.swLabel.Set(string(state))
}

func (bd data) call() {
	if bd.callback != nil {
		bd.callback()
	}
}

func (bd *data) getSw() *render.Switch {
	if bd.sw != nil {
		return bd.sw
	}
	bd.sw = render.NewSwitch(string(bd.state), map[string]render.Modifiable{
		string(Up):       render.NewColorBox(bd.size.X(), bd.size.Y(), bd.colors[Up]),
		string(Down):     render.NewColorBox(bd.size.X(), bd.size.Y(), bd.colors[Down]),
		string(Disabled): render.NewColorBox(bd.size.X(), bd.size.Y(), bd.colors[Disabled]),
	})
	return bd.sw
}

func (bd *data) textArgs(text string, font *render.Font) (str string, x float64, y float64) {
	textRect := layout2d.AlignRect(
		layout2d.HCenter|layout2d.VCenter,
		bd.bounds(),
		utils.TextMeasureRect(text, font),
		0,
	)
	return text, textRect.Min.X(), textRect.Min.Y()
}

func (bd *data) getSwLabel() *render.Switch {
	if bd.swLabel != nil {
		return bd.swLabel
	}
	bd.swLabel = render.NewSwitch(string(bd.state), map[string]render.Modifiable{
		string(Up):       bd.fonts[Up].NewText(bd.textArgs(bd.text, bd.fonts[Up])).ToSprite(),
		string(Down):     bd.fonts[Down].NewText(bd.textArgs(bd.text, bd.fonts[Down])).ToSprite(),
		string(Disabled): bd.fonts[Disabled].NewText(bd.textArgs(bd.text, bd.fonts[Disabled])).ToSprite(),
	})
	return bd.swLabel
}
