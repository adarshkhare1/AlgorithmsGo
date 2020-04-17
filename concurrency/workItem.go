package concurrency

import "sync"

// Task encapsulates a work item that should go in a work pool.
type WorkItem struct {
	// Err holds an error that occurred when running a workItem. Its
	Err error
	f func() error
}

// NewTask initializes a new task based on a given work
func NewWorkItem(f func() error) *WorkItem {
	return &WorkItem{f: f}
}

// Run runs a Task and does appropriate accounting via a
// given sync.WorkGroup.
func (t *WorkItem) Run(wg *sync.WaitGroup) {
	t.Err = t.f()
	wg.Done()
}