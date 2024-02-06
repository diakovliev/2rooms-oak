package button

import (
	"image/color"

	"github.com/diakovliev/2rooms-oak/packages/utils"
	"github.com/diakovliev/oak/v4/alg/intgeom"
	"github.com/diakovliev/oak/v4/render"
	"golang.org/x/image/colornames"
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
	ret = map[State]*render.Font{}
	for k, f := range fonts {
		ret[k] = utils.ColoredFont(f, fontColors[k])
	}
	return
}

func defaultSize() intgeom.Point2 {
	return intgeom.Point2{100, 100}
}
