package searching

import (
	"AlgorithmsGo/testHelper"
	"testing"
)

func TestFindMaxSubArrayBasic(t *testing.T) {
	test := MaxSubArrayBasic()
	funcName(t, test, FindMaxSubArray(test.Input))
}

func TestFindMaxSubArrayFullLength(t *testing.T) {
	test := MaxSubArrayFullLength()
	funcName(t, test, FindMaxSubArray(test.Input))
}

func TestFindMaxSubArrayEmpty(t *testing.T) {
	test := MaxSubArrayEmpty()
	funcName(t, test, FindMaxSubArray(test.Input))
}

func TestFindMaxSubArraySingleElement(t *testing.T) {
	test := MaxSubArraySingleElement()
	funcName(t, test, FindMaxSubArray(test.Input))
}

func TestFindMaxSubArrayAllNegative(t *testing.T) {
	test := MaxSubArrayAllNegative()
	funcName(t, test, FindMaxSubArray(test.Input))
}

func funcName(t *testing.T, test *MaxSubArrayTestInput, result SubArray, ) {
	testHelper.VerifyIntsAreEqual(t, test.outStart, result.Start)
	testHelper.VerifyIntsAreEqual(t, test.outEnd, result.End)
	testHelper.VerifyIntsAreEqual(t, test.outSum, result.Sum)
	//fmt.Printf("result start: %d end:%d sum:%d\n", result.Start, result.End, result.Sum)
}
