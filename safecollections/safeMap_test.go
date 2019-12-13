package safecollections

import (
	"Algorithms/testHelper"
	"testing"
)

func TestSafeMap(t *testing.T) {
	m := newSafeMap(5)
	m.Set("t0", "val0")
	m.Set("t1", "val1")
	s, _ := m.Get("t0")
	testHelper.VerifyStringsAreEqual(t, "val0", s)
	i := 0
	for item := range m.Iter() {
		switch item.Key {
		case "t0":
			testHelper.VerifyStringsAreEqual(t, "val0", item.Value)
		case "t1":
			testHelper.VerifyStringsAreEqual(t, "val1", item.Value)
		}
		i++
	}

}
