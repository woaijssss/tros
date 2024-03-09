package timer

import (
	"time"
)

type TrTimer interface {
	Reset(timeout time.Duration) bool
	Stop() bool
}

type trTimer struct {
	*time.Timer
}

func NewTimer(timeout time.Duration, workFunc func()) TrTimer {
	t := time.AfterFunc(timeout, workFunc)

	return t
}

func (tr *trTimer) Reset(timeout time.Duration) bool {
	return tr.Reset(timeout)
}

func (tr *trTimer) Stop() bool {
	return tr.Stop()
}
