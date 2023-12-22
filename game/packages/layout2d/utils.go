package layout2d

import (
	"github.com/diakovliev/oak/v4/alg/floatgeom"
)

type Rect = floatgeom.Rect2

func VAlignRect(alignment Alignment, parent Rect, rect Rect, margin float64) (ret Rect) {
	pw := parent.W()
	top := rect.Min.Y()
	left := parent.Min.X()
	w := rect.W()
	h := rect.H()
	oldMin := rect.Min
	newMin := oldMin
	switch {
	case alignment&Left == Left:
		newMin = floatgeom.Point2{left + margin, top}
	case alignment&VCenter == VCenter:
		newMin = floatgeom.Point2{left + (pw-w)/2, top}
	case alignment&Right == Right:
		newMin = floatgeom.Point2{left + pw - w - margin, top}
	}
	ret = Rect{
		Min: newMin,
		Max: floatgeom.Point2{
			newMin.X() + w,
			newMin.Y() + h,
		},
	}
	return
}

func HAlignRect(alignment Alignment, parent Rect, rect Rect, margin float64) (ret Rect) {
	ph := parent.H()
	top := parent.Min.Y()
	left := rect.Min.X()
	w := rect.W()
	h := rect.H()
	oldMin := rect.Min
	newMin := oldMin
	switch {
	case alignment&Top == Top:
		newMin = floatgeom.Point2{left, top + margin}
	case alignment&HCenter == HCenter:
		newMin = floatgeom.Point2{left, top + (ph-h)/2}
	case alignment&Bottom == Bottom:
		newMin = floatgeom.Point2{left, top + ph - h - margin}
	}
	ret = Rect{
		Min: newMin,
		Max: floatgeom.Point2{
			newMin.X() + w,
			newMin.Y() + h,
		},
	}
	return
}

func AlignRect(alignment Alignment, parent Rect, rect Rect, margin float64) (ret Rect) {
	ret = VAlignRect(alignment, parent, rect, margin)
	ret = HAlignRect(alignment, parent, ret, margin)
	return
}
