package safecollections

import "sync"

// Map type that can be safely shared between
// goroutines that require read/write access to a map
type safeMap struct {
	sync.RWMutex
	items map[string]string
}

func newSafeMap(size int) *safeMap {
	m := make (map[string]string, size)
	mutex := sync.RWMutex{}
	return &safeMap{mutex, m}
}

// Concurrent map item
type safeMapItem struct {
	Key   string
	Value string
}

// Sets a key in a concurrent map
func (sMap *safeMap) Set(key string, value string ){
	sMap.Lock()
	defer sMap.Unlock()

	sMap.items[key] = value
}

// Gets a key from a concurrent map
func (sMap *safeMap) Get(key string) (string, bool) {
	sMap.Lock()
	defer sMap.Unlock()

	value, ok := sMap.items[key]

	return value, ok
}

// Iterates over the items in a concurrent map
// Each item is sent over a channel, so that
// we can iterate over the map using the builtin range keyword
func (sMap *safeMap) Iter() <-chan safeMapItem {
	mapItem := make(chan safeMapItem)

	f := func() {
		sMap.Lock()
		defer sMap.Unlock()

		for k, v := range sMap.items {
			mapItem <- safeMapItem{k, v}
		}
		close(mapItem)
	}
	go f()

	return mapItem
}