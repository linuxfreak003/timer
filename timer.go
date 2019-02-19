package timer

import (
	"time"
)

type Timer struct {
	*time.Timer
	Start time.Time
	End   time.Time
}

func NewTimer(d time.Duration) *Timer {
	return &Timer{
		Timer: time.NewTimer(d),
		Start: time.Now(),
		End:   time.Now().Add(d),
	}
}

func (t *Timer) Reset(d time.Duration) bool {
	t.Start = time.Now()
	t.End = time.Now().Add(d)
	return t.Timer.Reset(d)
}

func (t *Timer) Remaining() time.Duration {
	return t.End.Sub(time.Now())
}

func (t *Timer) Add(d time.Duration) {
	t.End = t.End.Add(d)
	t.Timer = time.NewTimer(t.End.Sub(time.Now()))
}

func (t *Timer) Subtract(d time.Duration) {
	t.End = t.End.Add(d * -1)
	t.Timer = time.NewTimer(t.End.Sub(time.Now()))
}
