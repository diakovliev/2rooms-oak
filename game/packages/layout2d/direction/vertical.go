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
			// vectors
			func(layout *Layout, alignment layout2d.Alignment) (ret []common.Vector) {
				top := layout.pos.Y() + layout.margin
				for _, entity := range layout.entities {
					w := entity.W()
					h := entity.H()
					oldPos := floatgeom.Point2{entity.X(), entity.Y()}
					newPos := oldPos
					switch {
					case alignment&layout2d.Left == layout2d.Left:
						newPos = floatgeom.Point2{layout.pos.X() + layout.margin, top}
					case alignment&layout2d.VCenter == layout2d.VCenter:
						newPos = floatgeom.Point2{layout.pos.X() + layout.w/2 - w/2, top}
					case alignment&layout2d.Right == layout2d.Right:
						newPos = floatgeom.Point2{layout.pos.X() + layout.w - w - layout.margin, top}
					}
					ret = append(ret, common.Vector{
						Entity: entity,
						Delta:  newPos.Sub(oldPos),
						Old:    oldPos,
						New:    newPos,
						// TODO: get entity speed
						Speed: layout.speed,
					})
					top += h + layout.margin
				}
				return
			},
			// add
			func(layout *Layout, entities []common.Entity) {
				for _, entity := range entities {
					w := entity.W()
					h := entity.H()
					if float64(w)+2*layout.margin > layout.w {
						layout.w = float64(w) + 2*layout.margin
					}
					layout.h += float64(h) + layout.margin
					layout.entities = append(layout.entities, entity)
				}
			},
		),
	}
}
