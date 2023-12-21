package stimer

import (
	"sync"
	"time"

	"github.com/oakmound/oak/v4/event"

	oakscene "github.com/oakmound/oak/v4/scene"
)

type STimer struct {
	sync.Mutex
	duration time.Duration
	lastCall time.Time
	callback func()
}

func New(ctx *oakscene.Context, duration time.Duration, callback func()) (ret *STimer) {
	ret = &STimer{
		duration: duration,
		lastCall: time.Time{},
		callback: callback,
	}
	event.GlobalBind(ctx, event.Enter, ret.onFrame)
	return
}

func (st *STimer) onFrame(ev event.EnterPayload) event.Response {
	st.Lock()
	defer st.Unlock()
	if time.Since(st.lastCall) > st.duration {
		st.callback()
		st.lastCall = time.Now()
	}
	return event.ResponseNone
}
