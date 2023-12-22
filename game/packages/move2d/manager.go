package move2d

import (
	"sync"

	"github.com/diakovliev/2rooms-oak/packages/common"
	"github.com/diakovliev/oak/v4/alg/floatgeom"
	"github.com/diakovliev/oak/v4/event"

	oakscene "github.com/diakovliev/oak/v4/scene"
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
	newMove := newLinear(m, vector, speed)
	move, ok := m.moves[vector.Entity.CID()]
	if ok {
		// Replace move
		newMove.SetPos(move.Pos())
	}
	m.moves[newMove.CID()] = newMove
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
