package asyncwait

import (
	"context"
	"time"
)

// AsyncWait async wait representation
type AsyncWait interface {
	// Wait wait method
	Wait(func() bool) bool
}

var _ AsyncWait = (*asyncWait)(nil)

type asyncWait struct {
	tickerChan <-chan time.Time
	timeout    time.Duration
	doneCh     chan struct{}
}

// NewAsyncWait constructor for AsyncWait
func NewAsyncWait(timeout, pollInterval time.Duration) AsyncWait {
	return &asyncWait{
		tickerChan: time.NewTicker(pollInterval).C,
		timeout:    timeout,
		doneCh:     make(chan struct{}),
	}
}

// Wait while timeout, make polls every pollInterval for the predicate while is not truth
func (aw asyncWait) Wait(predicate func() bool) bool {
	ctx, cancel := context.WithTimeout(context.Background(), aw.timeout)
	defer cancel()

	for {
		select {
		case <-aw.doneCh:
			return true
		case <-ctx.Done():
			return false
		case <-aw.tickerChan:
			go func() {
				if predicate() {
					aw.doneCh <- struct{}{}
				}
			}()
		}
	}
}
