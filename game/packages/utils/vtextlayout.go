package utils

import (
	"fmt"

	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/render"
)

type VTextLayout struct {
	// Texts to render
	strings []fmt.Stringer
	// Used font
	f *render.Font
	// Margin between text
	margin float64
	// Calculated width and Height
	W, H float64
}

func NewVTextLayout(fi *render.Font, margin float64) *VTextLayout {
	var f *render.Font
	if f == nil {
		f = render.DefaultFont()
	} else {
		f = fi
	}
	return &VTextLayout{
		f:      f.Copy(),
		margin: margin,
		W:      margin,
		H:      margin,
	}
}

func (vtl VTextLayout) Add(strs ...fmt.Stringer) *VTextLayout {
	for _, str := range strs {
		w := vtl.f.MeasureString(str.String()).Ceil()
		h := vtl.f.Height()
		if float64(w)+2.*vtl.margin > vtl.W {
			vtl.W = float64(w) + 2*vtl.margin
		}
		vtl.H += float64(h) + vtl.margin
		vtl.strings = append(vtl.strings, str)
	}
	return &vtl
}

func (vtl VTextLayout) GetDims() (w, h float64) {
	return vtl.W, vtl.H
}

func (vtl VTextLayout) GetFDims() floatgeom.Point2 {
	return floatgeom.Point2{vtl.W, vtl.H}
}

func (vtl VTextLayout) Renderable() render.Renderable {
	renderables := make([]render.Renderable, 0, len(vtl.strings))
	top := vtl.margin
	for _, str := range vtl.strings {
		h := vtl.f.Height()
		renderables = append(
			renderables,
			vtl.f.NewStringerText(str, vtl.margin, float64(top)),
		)
		top += float64(h) + vtl.margin
	}
	return render.NewCompositeR(renderables...)
}
