package concurrency
import (
	"sync"
)

// Pool is a worker group that runs a number of workItem.
type WorkerPool struct {
	Items []*WorkItem

	concurrency int
	tasksChan   chan *WorkItem
	wg          sync.WaitGroup
}

// NewWorkerPool initializes a new pool with the given workItems and
// at the given concurrency.
func NewWorkerPool(workItems []*WorkItem, concurrency int) *WorkerPool {
	return &WorkerPool{
		Items:       workItems,
		concurrency: concurrency,
		tasksChan:   make(chan *WorkItem),
	}
}

// Run runs all work within the pool and blocks until it's
// finished.
func (p *WorkerPool) Run() {
	for i := 0; i < p.concurrency; i++ {
		go p.work()
	}
	p.wg.Add(len(p.Items))
	for _, task := range p.Items {
		p.tasksChan <- task
	}

	// all workers return
	close(p.tasksChan)

	p.wg.Wait()
}

// The work loop for any single goroutine.
func (p *WorkerPool) work() {
	for task := range p.tasksChan {
		task.Run(&p.wg)
	}
}
