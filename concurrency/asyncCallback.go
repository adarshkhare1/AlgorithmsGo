package concurrency

// Task encapsulates a work item that should go in a work pool.
type AsyncCallback struct {
	// Err holds an error that occurred when running a workItem. Its
	Err error
	callback func() error
	// result is only meaningful after Run has been called.
	result interface{}
}
