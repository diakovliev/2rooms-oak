package utils

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
)

type VAlignment int

type Entity interface {
	W() float64
	H() float64
	SetPos(p floatgeom.Point2)
}

const (
	VLeft VAlignment = iota
	VCenter
	VRight
)

type VLayout struct {
	entities  []Entity
	pos       floatgeom.Point2
	alignment VAlignment
	margin    float64
	w, h      float64
}

func NewVLayout(pos floatgeom.Point2, margin float64) VLayout {
	return VLayout{
		pos:    pos,
		margin: margin,
		w:      margin,
		h:      margin,
	}
}

func (vl VLayout) Add(e ...Entity) *VLayout {
	for _, ee := range e {
		w := ee.W()
		h := ee.H()
		if float64(w)+2*vl.margin > vl.w {
			vl.w = float64(w) + 2*vl.margin
		}
		vl.h += float64(h) + vl.margin
		vl.entities = append(vl.entities, ee)
	}
	return &vl
}

func (vl VLayout) GetDims() (w, h float64) {
	return vl.w, vl.h
}

func (vl VLayout) GetFDims() floatgeom.Point2 {
	return floatgeom.Point2{vl.w, vl.h}
}

func (vl VLayout) W() float64 {
	return vl.w
}

func (vl VLayout) H() float64 {
	return vl.h
}

// SetPos sets the position of the VLayout to the specified point.
//
// p: The new position for the VLayout.
func (vl *VLayout) SetPos(p floatgeom.Point2) {
	vl.pos = p
	vl.Rearrange(vl.alignment)
}

// Rearrange rearranges the entities in a VLayout according to the specified alignment.
//
// alignment is the vertical alignment to use.
func (vl *VLayout) Rearrange(alignment VAlignment) {
	vl.alignment = alignment
	top := vl.pos.Y() + vl.margin
	for _, ee := range vl.entities {
		w := ee.W()
		h := ee.H()
		var pos floatgeom.Point2
		switch vl.alignment {
		case VLeft:
			pos = floatgeom.Point2{vl.pos.X() + vl.margin, top}
		case VCenter:
			pos = floatgeom.Point2{vl.pos.X() + vl.w/2 - w/2, top}
		case VRight:
			pos = floatgeom.Point2{vl.pos.X() + vl.w - w - vl.margin, top}
		}
		ee.SetPos(pos)
		top += h + vl.margin
	}
}
