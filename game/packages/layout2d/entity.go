package layout2d

import "github.com/oakmound/oak/v4/alg/floatgeom"

// Entity is an entity that can be placed in a layout.
type Entity interface {
	X() float64
	Y() float64
	W() float64
	H() float64
	SetPos(p floatgeom.Point2)
}
