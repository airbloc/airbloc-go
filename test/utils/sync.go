package testutils

import (
	"errors"
	"sync"
	"time"
)

var (
	ErrTimeout = errors.New("waitgroup timeout")
)

// TimeoutWaitGroup is a WaitGroup with timeout support.
type TimeoutWaitGroup struct {
	timeout time.Duration
	sync.WaitGroup
}

// NewTimeoutWaitGroup creates new WaitGroup with timeout.
func NewTimeoutWaitGroup(timeout time.Duration) *TimeoutWaitGroup {
	return &TimeoutWaitGroup{timeout: timeout}
}

// Wait blocks until the WaitGroup counter is zero.
// An ErrTimeout is returned if it exceeds timeout.
func (twg *TimeoutWaitGroup) Wait() error {
	timeout := time.After(twg.timeout)
	finished := make(chan bool)

	go func() {
		twg.WaitGroup.Wait()
		finished <- true
	}()

	select {
	case <-timeout:
		return ErrTimeout
	case <-finished:
		return nil
	}
}
