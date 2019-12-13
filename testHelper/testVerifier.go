package testHelper

import "testing"

func VerifyStringsAreEqual(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("fail actual: %q, expected: %q", actual, expected)
	}
}

func VerifyIntsAreEqual(t *testing.T, expected int, actual int) {
		if actual != expected {
			t.Errorf("fail actual: %d, expected: %d", actual, expected)
		}
}

func VerifyUintsAreEqual(t *testing.T, expected uint, actual uint) {
	if actual != expected {
		t.Errorf("fail actual: %d, expected: %d", actual, expected)
	}
}

func VerifyErrorNotNil(t *testing.T, err error) {
	if err == nil {
		t.Errorf("fail expecting error, error is nil")
	}
}

func VerifyErrorNil(t *testing.T, err error) {
	if err != nil {
		t.Log(err)
		t.Errorf("fail expecting nil error, error is not nil")
	}
}

func VerifyArraysAreEqual(t *testing.T, expected, actual []int) {
	if len(expected) != len(actual) {
		t.Errorf("fail array size mismatch actual: %d, expected: %d", len(actual), len(expected))
	}
	for i, v := range actual {
		if v != expected[i] {
			t.Errorf("fail array value mismatch on %d, actual: %d, expected: %d", i, v, expected[i])
		}
	}
}