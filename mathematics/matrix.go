package mathematics

import (
	"errors"
	"fmt"
)

type Matrix struct {
	Value   [][]int
	numRows    int
	numColumns int
}

func NewMatrix(nRow int, nColumn int) *Matrix {
	val := make([][]int, nRow)
	for i:=0;i < nRow ; i++ {
		val[i] = make([]int, nColumn)
	}
	return &Matrix{numRows: nRow, numColumns: nColumn, Value: val}
}

// UpdateRow updates the row in the matrix, return error if row can't be added
func (m *Matrix) UpdateRow(newRow []int, rowNum int) error{
	if len(newRow) != m.numColumns{
		return errors.New("invalid size row")
	}
	m.Value[rowNum] = newRow
	return nil
}

//Multiply Matrix m with m2 and return the result
func (m *Matrix)Multiply(m2 *Matrix) *Matrix{
	if m.numColumns != m2.numRows {
		return nil
	}
	result := NewMatrix(m.numRows,m2.numColumns)
	for i:=0;i < m.numRows ; i++{
		_ = result.UpdateRow(make([]int, m.numRows), i)
		for j := 0; j < m2.numColumns; j++ {
			result.Value[i][j] = 0
			for k:= 0; k < m.numColumns; k++{
				result.Value[i][j] += m.Value[i][k]*m2.Value[k][j]
			}
		}
	}
	return result
}

//Add Matrix m with m2 and return the result
func (m *Matrix)Add(m2 *Matrix) *Matrix {
	if m.numColumns != m2.numColumns || m.numRows != m2.numRows {
		return nil
	}
	result := NewMatrix(m.numRows,m.numColumns)
	for i:=0;i < m.numRows ; i++ {
		_ = result.UpdateRow(make([]int, m.numRows), i)
		for j := 0; j < m2.numColumns; j++ {
			result.Value[i][j] =  m.Value[i][j]+m2.Value[i][j]
		}
	}
	return result
}

//Subtract Matrix m with m2 and return the result
func (m *Matrix)Subtract(m2 *Matrix) *Matrix {
	if m.numColumns != m2.numColumns || m.numRows != m2.numRows {
		return nil
	}
	result := NewMatrix(m.numRows,m.numColumns)
	for i:=0;i < m.numRows ; i++ {
		_ = result.UpdateRow(make([]int, m.numRows), i)
		for j := 0; j < m2.numColumns; j++ {
			result.Value[i][j] =  m.Value[i][j]-m2.Value[i][j]
		}
	}
	return result
}

func (m *Matrix)Transpose()  *Matrix {
	result := NewMatrix(m.numColumns,m.numRows)
	for i := 0; i < m.numColumns; i++ {
		_ = result.UpdateRow(make([]int, m.numRows), i)
		for j := 0; j < m.numRows; j++ {
			result.Value[i][j] = m.Value[j][i]
		}
	}
	return result
}
//Print the rows of matrix, one row in a line
func (m *Matrix)Print() {
	for i := 0; i < len(m.Value); i++ {
		fmt.Printf("%v\n", m.Value[i]) // each one prints 5
	}
}