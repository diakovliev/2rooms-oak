package move2d

import (
	"sync"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/event"
)

type LinearMove struct {
	sync.Mutex
	// Parent
	parent *Manager
	// vector
	vector common.Vector
	// Current position
	pos floatgeom.Point2
	// Current speed
	speed floatgeom.Point2
}

func newLinear(m *Manager, vector common.Vector, speed floatgeom.Point2) *LinearMove {
	return &LinearMove{
		parent: m,
		vector: vector,
		pos:    vector.Old,
		speed:  speed,
	}
}

// Returns the parent
func (lm *LinearMove) Parent() *Manager {
	lm.Lock()
	defer lm.Unlock()
	return lm.parent
}

// Returns the CID of the move
func (lm *LinearMove) CID() event.CallerID {
	lm.Lock()
	defer lm.Unlock()
	return lm.vector.Entity.CID()
}

// Returns the vector of the move
func (lm *LinearMove) Vector() common.Vector {
	lm.Lock()
	defer lm.Unlock()
	return lm.vector
}

// Returns the current position of the move
func (lm *LinearMove) Pos() floatgeom.Point2 {
	lm.Lock()
	defer lm.Unlock()
	return lm.pos
}

func (lm *LinearMove) SetPos(pos floatgeom.Point2) {
	lm.Lock()
	defer lm.Unlock()
	lm.pos = pos
	lm.vector.Entity.SetPos(lm.pos)
}

// Returns the speed of the move
func (lm *LinearMove) Speed() floatgeom.Point2 {
	lm.Lock()
	defer lm.Unlock()
	return lm.speed
}

// Returns true if the move is complete
func (lm *LinearMove) IsComplete() bool {
	lm.Lock()
	defer lm.Unlock()
	return lm.pos == lm.vector.New
}

// Do performs the move
func (lm *LinearMove) Do(ev event.EnterPayload) bool {
	lm.Lock()
	defer lm.Unlock()
	speed := lm.speed.Magnitude()
	delta := speed * ev.TickPercent
	if lm.pos.Distance(lm.vector.New) < delta {
		lm.pos = lm.vector.New
	} else {
		lm.pos = lm.pos.Add(lm.vector.New.Sub(lm.pos).Normalize().MulConst(delta))
		// lm.pos = lm.pos.Add(lm.vector.New.Sub(lm.pos).Normalize().Mul(lm.speed))
	}
	lm.vector.Entity.SetPos(lm.pos)
	// IsComplete()
	return lm.pos == lm.vector.New
}
