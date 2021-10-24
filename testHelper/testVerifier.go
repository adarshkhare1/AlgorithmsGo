package testHelper

import (
	"reflect"
	"testing"
)

func VerifyTrue(t *testing.T, actual bool) {
	if !actual {
		t.Errorf("Condition is false, expected: true")
	}
}

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

func VerifyNil(t *testing.T, i interface{}) {
	if !reflect.ValueOf(i).IsNil() {
		t.Errorf("fail expecting nil, value is not nil")
	}
}

func VerifyNotNil(t *testing.T, i interface{}) {
	if reflect.ValueOf(i).IsNil() {
		t.Errorf("fail expecting not nil, value is nil")
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

func VerifyArrayIsSortedAscending(t *testing.T, nums []int) {
	for i := 1; i < len(nums); i++ {
		VerifyTrue(t, nums[i-1] <= nums[i])
	}
}
