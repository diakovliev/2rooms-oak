package move2d

import (
	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/event"
)

type Move interface {
	// Returns the parent
	Parent() *Manager
	// Returns the CID of the move
	CID() event.CallerID
	// Returns the vector of the move
	Vector() common.Vector
	// Returns the current position of the move
	Pos() floatgeom.Point2
	// Sets the position of the move
	SetPos(floatgeom.Point2)
	// Returns the speed of the move
	Speed() floatgeom.Point2
	// Returns true if the move is complete
	IsComplete() bool
	// Do performs the move
	Do(ev event.EnterPayload) bool
}
