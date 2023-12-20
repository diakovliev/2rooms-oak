package direction

import (
	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	oakscene "github.com/oakmound/oak/v4/scene"
)

// HLayout is a Layout with a horizontal layout.
type HLayout struct {
	*Layout
}

// Horizontal returns a GLayout with a horizontal layout.
//
// It takes in the position of the layout (pos floatgeom.Point2) and the margin (margin float64).
// It returns a HLayout.
func Horizontal(
	ctx *oakscene.Context,
	pos floatgeom.Point2,
	speed floatgeom.Point2,
	margin float64,
) HLayout {
	return HLayout{
		Layout: newLayout(
			ctx,
			pos,
			speed,
			margin,
			func(hl *Layout, alignment layout2d.Alignment) (ret []common.Vector) {
				left := hl.pos.X() + hl.margin
				for _, ee := range hl.entities {
					w := ee.W()
					h := ee.H()
					oldPos := floatgeom.Point2{ee.X(), ee.Y()}
					newPos := oldPos
					switch {
					case alignment&layout2d.Top == layout2d.Top:
						newPos = floatgeom.Point2{left, hl.pos.Y() + hl.margin}
					case alignment&layout2d.HCenter == layout2d.HCenter:
						newPos = floatgeom.Point2{left, hl.pos.Y() + (hl.h-h)/2}
					case alignment&layout2d.Bottom == layout2d.Bottom:
						newPos = floatgeom.Point2{left, hl.pos.Y() + hl.h - h - hl.margin}
					}
					ret = append(ret, common.Vector{
						Entity: ee,
						Delta:  newPos.Sub(oldPos),
						Old:    oldPos,
						New:    newPos,
						// TODO: get entity speed
						Speed: hl.speed,
					})
					left += w + hl.margin
				}
				return
			},
			func(hl *Layout, e []common.Entity) {
				for _, ee := range e {
					w := ee.W()
					h := ee.H()
					if float64(h)+2*hl.margin > hl.h {
						hl.h = float64(h) + 2*hl.margin
					}
					hl.w += float64(w) + hl.margin
					hl.entities = append(hl.entities, ee)
				}
			},
		),
	}
}
