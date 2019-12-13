package safecollections

import (
	"Algorithms/testHelper"
	"testing"
)

func TestSafeList(t *testing.T) {
	l := newSafeListSlice()
	l.Append("test1")
	l.Append("test2")
	testHelper.VerifyIntsAreEqual(t, 2, len(l.items))
	i := 0
	for val := range l.Iter() {
		switch i {
		case 0:
			testHelper.VerifyStringsAreEqual(t, "test1", val.Value)
		case 1:
			testHelper.VerifyStringsAreEqual(t, "test2", val.Value)
		}
		i++
	}
}