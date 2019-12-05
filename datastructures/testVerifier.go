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

func verifyErrorNotNil(t *testing.T, err error) {
	if err == nil {
		t.Errorf("fail expecting error, error is nil")
	}
}

func verifyErrorNil(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.Errorf("fail expecting nil error, error is not nil")
	}
}


