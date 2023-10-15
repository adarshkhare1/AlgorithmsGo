package mathematics

import (
	"AlgorithmsGo/testHelper"
	"testing"
)

func TestFindGreatestCommonDivisor(t *testing.T) {
	result := FindGreatestCommonDivisor(4, 6)
	testHelper.VerifyIntsAreEqual(t, 2, result)
	result = FindGreatestCommonDivisor(6, 4)
	testHelper.VerifyIntsAreEqual(t, 2, result)
	result = FindGreatestCommonDivisor(3, 5)
	testHelper.VerifyIntsAreEqual(t, 1, result)
	result = FindGreatestCommonDivisor(6, 6)
	testHelper.VerifyIntsAreEqual(t, 6, result)
}
