package move2d

import (
	"sync"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/oakmound/oak/v4/alg/floatgeom"
	"github.com/oakmound/oak/v4/event"

	oakscene "github.com/oakmound/oak/v4/scene"
)

type Manager struct {
	sync.Mutex
	moves map[event.CallerID]Move
}

func New(ctx *oakscene.Context) (ret *Manager) {
	ret = &Manager{
		moves: make(map[event.CallerID]Move),
	}
	event.GlobalBind(ctx, event.Enter, ret.onFrame)
	return
}

func (m *Manager) Start(vector common.Vector, speed floatgeom.Point2) (move Move) {
	m.Lock()
	defer m.Unlock()
	move, ok := m.moves[vector.Entity.CID()]
	if !ok {
		// Create new move
		move = newLinear(m, vector, speed)
	} else {
		// Replace move
		pos := move.Pos()
		move = newLinear(m, vector, speed)
		move.SetPos(pos)
	}
	m.moves[move.CID()] = move
	return
}

func (m *Manager) onFrame(ev event.EnterPayload) event.Response {
	m.Lock()
	defer m.Unlock()
	for _, move := range m.moves {
		if move.Do(ev) {
			delete(m.moves, move.CID())
		}
	}
	return event.ResponseNone
}
