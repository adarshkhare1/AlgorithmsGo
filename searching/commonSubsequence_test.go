package searching

import (
	"AlgorithmsGo/testHelper"
	"testing"
)

func TestLongestCommonSubsequence(t *testing.T) {
	x := []byte("BDCABAC")
	y := []byte("ABCBDAB")
	length, seq := LongestCommonSubsequence(x, y)
	testHelper.VerifyIntsAreEqual(t, LCSLength(x, y), length)
	testHelper.VerifyStringsAreEqual(t, "BDAB", string(seq))
}

func TestLCSPerfectMatch(t *testing.T) {
	x := []byte("ABC")
	y := []byte("ABC")
	length, seq := LongestCommonSubsequence(x, y)
	testHelper.VerifyIntsAreEqual(t, LCSLength(x, y), length)
	testHelper.VerifyStringsAreEqual(t, "ABC", string(seq))
}
