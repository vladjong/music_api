package timer

import (
	"time"
)

type State int

const (
	StateIdle State = iota
	StateActive
	StateExpired
)

type Timer struct {
	c         chan time.Time
	C         chan time.Time
	duration  time.Duration
	State     State
	fn        func()
	startedAt time.Time
	t         *time.Timer
}

func NewTimer(d time.Duration) *Timer {
	c := make(chan time.Time, 1)
	t := new(Timer)
	t.c = c
	t.C = c
	t.duration = d
	t.fn = func() {
		t.State = StateExpired
		t.c <- time.Now()
	}
	return t
}

func (t *Timer) Pause() bool {
	if t.State != StateActive {
		return false
	}
	if !t.t.Stop() {
		t.State = StateExpired
		return false
	}
	t.State = StateIdle
	dur := time.Since(t.startedAt)
	t.duration = t.duration - dur
	return true
}

func (t *Timer) Start() bool {
	if t.State != StateIdle {
		return false
	}
	t.startedAt = time.Now()
	t.State = StateActive
	t.t = time.AfterFunc(t.duration, t.fn)
	return true
}

func (t *Timer) Stop() bool {
	if t.State != StateActive {
		return false
	}
	t.startedAt = time.Now()
	t.State = StateExpired
	t.t.Stop()
	return true
}
