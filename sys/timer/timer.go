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

func (xt *trTimer) Reset(timeout time.Duration) bool {
	return xt.Reset(timeout)
}

func (xt *trTimer) Stop() bool {
	return xt.Stop()
}
