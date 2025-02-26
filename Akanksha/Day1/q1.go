// Question 1

package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	Rows     int
	Cols     int
	Elements [][]int
}

func (m *Matrix) GetRows() int {
	return m.Rows
}

func (m *Matrix) GetCols() int {
	return m.Cols
}

func (m *Matrix) SetElement(i, j, value int) {
	m.Elements[i][j] = value
}

func (m *Matrix) AddMatrix(other Matrix) Matrix {
	var result Matrix
	result.Rows = m.Rows
	result.Cols = m.Cols
	result.Elements = make([][]int, m.Rows)
	for i := 0; i < m.Rows; i++ {
		result.Elements[i] = make([]int, m.Cols)
		for j := 0; j < m.Cols; j++ {
			result.Elements[i][j] = m.Elements[i][j] + other.Elements[i][j]
		}
	}
	return result
}

func (m *Matrix) ToJSON() string {
	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("Error converting to JSON")
		return ""
	}
	return string(jsonData)
}

func main() {
	m1 := Matrix{}
	m1.Rows = 2
	m1.Cols = 2
	m1.Elements = [][]int{{1, 2}, {3, 4}}

	m2 := Matrix{}
	m2.Rows = 2
	m2.Cols = 2
	m2.Elements = [][]int{{5, 6}, {7, 8}}

	result := m1.AddMatrix(m2)

	fmt.Println("Resulting Matrix in JSON:")
	fmt.Println(result.ToJSON())
}
