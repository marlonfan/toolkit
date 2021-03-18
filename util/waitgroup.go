package util

import (
	"sync"
)

// WaitGroupWrapper .
type WaitGroupWrapper struct {
	wg sync.WaitGroup
}

// AddAndRun .
func (w *WaitGroupWrapper) AddAndRun(cb func()) {
	w.wg.Add(1)
	go func() {
		cb()
		w.wg.Done()
	}()
}

// Wait .
func (w *WaitGroupWrapper) Wait() {
	w.wg.Wait()
}
