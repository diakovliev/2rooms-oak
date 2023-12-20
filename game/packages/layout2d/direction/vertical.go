package direction

import (
	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/2rooms-oak/packages/layout2d"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	oakscene "github.com/oakmound/oak/v4/scene"
)

// VLayout is a vertical layout
type VLayout struct {
	*Layout
}

// Vertical creates a vertical layout with the given position and margin.
//
// It takes the position of the layout as a floatgeom.Point2 and the margin as a float64.
// It returns a VLayout structure.
func Vertical(
	ctx *oakscene.Context,
	pos floatgeom.Point2,
	speed floatgeom.Point2,
	margin float64,
) VLayout {
	return VLayout{
		Layout: newLayout(
			ctx,
			pos,
			speed,
			margin,
			func(vl *Layout, alignment layout2d.Alignment) (ret []common.Vector) {
				top := vl.pos.Y() + vl.margin
				for _, ee := range vl.entities {
					w := ee.W()
					h := ee.H()
					oldPos := floatgeom.Point2{ee.X(), ee.Y()}
					newPos := oldPos
					switch {
					case alignment&layout2d.Left == layout2d.Left:
						newPos = floatgeom.Point2{vl.pos.X() + vl.margin, top}
					case alignment&layout2d.VCenter == layout2d.VCenter:
						newPos = floatgeom.Point2{vl.pos.X() + vl.w/2 - w/2, top}
					case alignment&layout2d.Right == layout2d.Right:
						newPos = floatgeom.Point2{vl.pos.X() + vl.w - w - vl.margin, top}
					}
					ret = append(ret, common.Vector{
						Entity: ee,
						Delta:  newPos.Sub(oldPos),
						Old:    oldPos,
						New:    newPos,
						// TODO: get entity speed
						Speed: vl.speed,
					})
					top += h + vl.margin
				}
				return
			},
			func(vl *Layout, e []common.Entity) {
				for _, ee := range e {
					w := ee.W()
					h := ee.H()
					if float64(w)+2*vl.margin > vl.w {
						vl.w = float64(w) + 2*vl.margin
					}
					vl.h += float64(h) + vl.margin
					vl.entities = append(vl.entities, ee)
				}
			},
		),
	}
}
