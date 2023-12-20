package common

import (
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/event"
)

// Entity is an entity that can be placed in a layout.
type Entity interface {
	CID() event.CallerID
	X() float64
	Y() float64
	W() float64
	H() float64
	SetPos(p floatgeom.Point2)
}
