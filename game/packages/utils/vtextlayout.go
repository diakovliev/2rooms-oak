package utils

import (
	"fmt"

	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/render"
)

type VTextLayout struct {
	// Texts to render
	strings []fmt.Stringer
	// Used font
	f *render.Font
	// Margin between text
	margin float64
	// Calculated width and Height
	w, h float64
}

func NewVTextLayout(f *render.Font, w, margin float64) *VTextLayout {
	if f == nil {
		f = render.DefaultFont()
	}
	return &VTextLayout{
		f:      f.Copy(),
		margin: margin,
		w:      w,
		h:      margin,
	}
}

func (vtl VTextLayout) W() float64 {
	return vtl.w
}

func (vtl VTextLayout) H() float64 {
	return vtl.h
}

func (vtl VTextLayout) maxW() float64 {
	return vtl.w - vtl.margin*2
}

func (vtl VTextLayout) Add(strs ...fmt.Stringer) *VTextLayout {
	for _, str := range strs {
		lines, rect, _, err := TextFitWidth(str.String(), vtl.maxW(), vtl.f)
		if err != nil {
			// TODO: handle error
			panic(err)
		}
		if rect.W() > vtl.maxW() {
			// TODO: handle error
			panic(fmt.Errorf("text too long"))
		}
		vtl.h += rect.H() + float64(len(lines))*vtl.margin
		vtl.strings = append(vtl.strings, str)
	}
	return &vtl
}

func (vtl VTextLayout) GetDims() (w, h float64) {
	return vtl.w, vtl.h
}

func (vtl VTextLayout) GetFDims() floatgeom.Point2 {
	return floatgeom.Point2{vtl.w, vtl.h}
}

func (vtl VTextLayout) Renderable() render.Renderable {
	renderables := make([]render.Renderable, 0, len(vtl.strings))
	top := vtl.margin
	for _, str := range vtl.strings {
		lines, rect, lh, err := TextFitWidth(str.String(), vtl.maxW(), vtl.f)
		if err != nil {
			// TODO: handle error
			panic(err)
		}
		if rect.W() > vtl.maxW() {
			// TODO: handle error
			panic(fmt.Errorf("text too long"))
		}
		if len(lines) == 1 {
			// Stringers must be a single liners
			renderables = append(
				renderables,
				vtl.f.NewStringerText(str, vtl.margin, float64(top)),
			)
			top += rect.H() + vtl.margin
			continue
		}
		for _, line := range lines {
			// Mutli line static text
			renderables = append(
				renderables,
				vtl.f.NewText(line, vtl.margin, float64(top)),
			)
			top += lh + vtl.margin
		}
	}
	return render.NewCompositeR(renderables...)
}
