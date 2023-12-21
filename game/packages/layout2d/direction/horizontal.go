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
			// vectors
			func(layout *Layout, alignment layout2d.Alignment) (ret []common.Vector) {
				left := layout.pos.X() + layout.margin
				for _, entity := range layout.entities {
					w := entity.W()
					h := entity.H()
					oldPos := floatgeom.Point2{entity.X(), entity.Y()}
					newPos := oldPos
					switch {
					case alignment&layout2d.Top == layout2d.Top:
						newPos = floatgeom.Point2{left, layout.pos.Y() + layout.margin}
					case alignment&layout2d.HCenter == layout2d.HCenter:
						newPos = floatgeom.Point2{left, layout.pos.Y() + (layout.h-h)/2}
					case alignment&layout2d.Bottom == layout2d.Bottom:
						newPos = floatgeom.Point2{left, layout.pos.Y() + layout.h - h - layout.margin}
					}
					ret = append(ret, common.Vector{
						Entity: entity,
						Delta:  newPos.Sub(oldPos),
						Old:    oldPos,
						New:    newPos,
						// TODO: get entity speed
						Speed: layout.speed,
					})
					left += w + layout.margin
				}
				return
			},
			// add
			func(layout *Layout, entity []common.Entity) {
				for _, ee := range entity {
					w := ee.W()
					h := ee.H()
					if float64(h)+2*layout.margin > layout.h {
						layout.h = float64(h) + 2*layout.margin
					}
					layout.w += float64(w) + layout.margin
					layout.entities = append(layout.entities, ee)
				}
			},
		),
	}
}
