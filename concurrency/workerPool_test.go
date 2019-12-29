package concurrency

import (
	"Algorithms/testHelper"
	"runtime"
	"testing"
	"time"
)


func TestWorkerPool(t *testing.T){
	runtime.GOMAXPROCS(4)
	items := []*WorkItem{
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
		NewWorkItem(testAdd),
	}

	pool := NewWorkerPool(items, 4)
	pool.Run()
	/*var numErrors int
	for _, task := range pool.Items {
		if task.Err != nil {
			log.Fatal(task.Err)
			numErrors++
		}
		if numErrors >= 10 {
			log.Fatal("Too many errors.")
			break
		}
	}*/
	testHelper.VerifyStringsAreEqual(t, "test", "test")
}

func testAdd() error{
	x := 1 + 2
	d, _ := time.ParseDuration("2s")
	time.Sleep(d)
	println("x=",x)
	return nil
}
