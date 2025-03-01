package main

import (
	"encoding/json"
	"fmt"
)

// Matrix struct represents a 2D matrix
// It contains the number of rows, number of columns, and a 2D slice to store the elements.
type Matrix struct {
	Rows     int     // Number of rows in the matrix
	Cols     int     // Number of columns in the matrix
	Elements [][]int // 2D slice to store matrix elements
}

// NewMatrix function initializes a matrix with the given rows and columns
func NewMatrix(rows, cols int) Matrix {
	elements := make([][]int, rows)
	for i := range elements {
		elements[i] = make([]int, cols)
	}
	return Matrix{Rows: rows, Cols: cols, Elements: elements}
}

// GetRows method returns the number of rows in the matrix
func (m *Matrix) GetRows() int {
	return m.Rows
}

// GetCols method returns the number of columns in the matrix
func (m *Matrix) GetCols() int {
	return m.Cols
}

// SetElement method sets the value of an element at a specific row and column
func (m *Matrix) SetElement(row, col, value int) {
	if row >= 0 && row < m.Rows && col >= 0 && col < m.Cols {
		m.Elements[row][col] = value
	} else {
		fmt.Println("Index out of bounds")
	}
}

// AddMatrix method adds two matrices and returns the result
func (m *Matrix) AddMatrix(other Matrix) Matrix {
	result := NewMatrix(m.Rows, m.Cols)
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

// PrintMatrixJSON method prints the matrix in JSON format
func (m *Matrix) PrintMatrixJSON() {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}
	fmt.Println(string(jsonData))
}

func main() {
	// Create a 3x3 matrix
	matrix1 := NewMatrix(3, 3)
	matrix2 := NewMatrix(3, 3)

	// Set some elements in the matrices
	matrix1.SetElement(0, 0, 1)
	matrix1.SetElement(1, 1, 5)
	matrix1.SetElement(2, 2, 9)

	matrix2.SetElement(0, 0, 2)
	matrix2.SetElement(1, 1, 3)
	matrix2.SetElement(2, 2, 4)

	// Add matrices
	resultMatrix := matrix1.AddMatrix(matrix2)

	// Print result matrix as JSON
	resultMatrix.PrintMatrixJSON()
}
