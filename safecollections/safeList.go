package safecollections

import "sync"

// Slice type that can be safely shared between goroutines
type safeListSlice struct {
	sync.RWMutex
	items []string
}

func newSafeListSlice() *safeListSlice {
	m := make ([]string, 0)
	mutex := sync.RWMutex{}
	return &safeListSlice{mutex, m}
}

// Concurrent slice item
type safeListItem struct {
	Index int
	Value string
}

// Appends an item to the concurrent slice
func (sList *safeListSlice) Append(item string) {
	sList.Lock()
	defer sList.Unlock()

	sList.items = append(sList.items, item)
}

// Iterates over the items in the concurrent slice
// Each item is sent over a channel, so that
// we can iterate over the slice using the builtin range keyword
func (sList *safeListSlice) Iter() <-chan safeListItem {
	sListItem := make(chan safeListItem)

	f := func() {
		sList.Lock()
		defer sList.Unlock()
		for index, value := range sList.items {
			sListItem <- safeListItem{index, value}
		}
		close(sListItem)
	}
	go f()

	return sListItem
}
