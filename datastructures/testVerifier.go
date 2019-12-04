package datastructures

import "testing"

func verifyString(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("fail actual: %q, expected: %q", actual, expected)
	}
}

func verifyInt(t *testing.T, expected int, actual int) {
		if actual != expected {
			t.Errorf("fail actual: %d, expected: %d", actual, expected)
		}
}
