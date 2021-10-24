package mathematics

import (
	"Algorithms/testHelper"
	"testing"
)

func TestMatrixBuild(t *testing.T) {
	m := NewMatrix(2,3)
	array2 := []int{1, 2, 3}
	array3 := []int{3, 3, 3, 3}
	array4 := []int{4, 5, 6}
	testHelper.VerifyErrorNil(t, m.UpdateRow(array2, 0))
	testHelper.VerifyErrorNotNil(t, m.UpdateRow(array3, 1)) // adding incorrect size row
	testHelper.VerifyErrorNil(t,m.UpdateRow(array4, 1))
	testHelper.VerifyIntsAreEqual(t, 2, len(m.Value))
	testHelper.VerifyArraysAreEqual(t, array2, m.Value[0])
	testHelper.VerifyArraysAreEqual(t, array4, m.Value[1])
	m.Print()
}
func TestMatrix_Transpose(t *testing.T) {
	testList := [][]int{{1, 2, 3}, {4, 5, 6}}
	testResult := [][]int{{1, 4}, {2, 5}, {3, 6}}
	m := buildTestMatrix(testList)
	m = m.Transpose()
	m.Print()
	testHelper.VerifyIntsAreEqual(t, len(testResult[0]), m.numColumns)
	testHelper.VerifyIntsAreEqual(t, len(testResult), m.numRows)
	verifyMatrixValue(t, testResult, m)

}

func TestMatrix_Add(t *testing.T) {
	testF := [][]int{{4, 7}, {1, 3}}
	testA := [][]int{{2, 4, 6}, {3, 5, 7}}
	testB := [][]int{{1, 3, 5}, {2, 4, 6}}
	testC := [][]int{{3, 7, 11}, {5, 9, 13}}
	mA := buildTestMatrix(testA)
	mB := buildTestMatrix(testB)
	testHelper.VerifyNil(t, mA.Add( buildTestMatrix(testF)))
	mC := mA.Add(mB)
	testHelper.VerifyIntsAreEqual(t, len(testC), mC.numRows)
	testHelper.VerifyIntsAreEqual(t, len(testC[0]), mC.numColumns)
	verifyMatrixValue(t,  testC, mC)
}

func TestMatrix_Subtract(t *testing.T) {
	testF := [][]int{{4, 7}, {1, 3}}
	testA := [][]int{{4, 7, 12}, {1, 3, 9}}
	testB := [][]int{{1, 3, 5}, {2, 4, 6}}
	testC := [][]int{{3, 4, 7}, {-1, -1, 3}}
	mA := buildTestMatrix(testA)
	mB := buildTestMatrix(testB)
	testHelper.VerifyNil(t, mA.Subtract( buildTestMatrix(testF)))
	mC := mA.Subtract(mB)
	testHelper.VerifyIntsAreEqual(t, len(testC), mC.numRows)
	testHelper.VerifyIntsAreEqual(t, len(testC[0]), mC.numColumns)
	verifyMatrixValue(t,  testC, mC)
}

func TestMatrix_Multiply(t *testing.T) {
	testA := [][]int{{4, 7, 12}, {1, -3, 9}}
	testB := [][]int{{1, 3, 5}, {2, 4, 6}, {4, 6, -1}}
	testC := [][]int{{66, 112, 50}, {31, 45, -22}}
	mA := buildTestMatrix(testA)
	mB := buildTestMatrix(testB)
	testHelper.VerifyNil(t, mB.Multiply(mA))
	mC := mA.Multiply(mB)
	testHelper.VerifyNotNil(t, mC)
	testHelper.VerifyIntsAreEqual(t, len(testC), mC.numRows)
	testHelper.VerifyIntsAreEqual(t, len(testC[0]), mC.numColumns)
	verifyMatrixValue(t,  testC, mC)
}

func buildTestMatrix(values [][]int) *Matrix{
	m := NewMatrix(len(values),len(values[0]))
	for i:= 0; i< len(values); i++ {
		m.UpdateRow(values[i], i)
	}
	return m
}

func verifyMatrixValue(t *testing.T, testResult [][]int, m *Matrix) {
	for i := 0; i < len(testResult); i++ {
		testHelper.VerifyArraysAreEqual(t, testResult[i], m.Value[i])
	}
}