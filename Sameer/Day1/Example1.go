package main

import "fmt"

type matrix struct {
	row    int
	col    int
	values [][]int
}

func getRows(m matrix) int {
	return m.row
}

func getCols(m matrix) int {
	return m.col
}

func setValue(m *matrix, i int, j int, val int) {
	m.values[i][j] = val
}

func addMat(m1 matrix, m2 matrix) matrix {
	var m3 matrix
	m3.row = m1.row
	m3.col = m1.col
	m3.values = make([][]int, m3.row)
	for i := 0; i < m3.row; i++ {
		m3.values[i] = make([]int, m3.col)
	}
	for i := 0; i < m3.row; i++ {
		for j := 0; j < m3.col; j++ {
			m3.values[i][j] = m1.values[i][j] + m2.values[i][j]
		}
	}
	return m3
}

func multiplyMat(m1 matrix, m2 matrix) matrix {
	var m3 matrix
	m3.row = m1.row
	m3.col = m2.col
	m3.values = make([][]int, m3.row)
	for i := 0; i < m3.row; i++ {
		m3.values[i] = make([]int, m3.col)
	}
	for i := 0; i < m3.row; i++ {
		for j := 0; j < m3.col; j++ {
			m3.values[i][j] = 0
			for k := 0; k < m1.col; k++ {
				m3.values[i][j] += m1.values[i][k] * m2.values[k][j]
			}
		}
	}
	return m3
}

func main() {

	matrix1 := matrix{2, 2, [][]int{{1, 2}, {3, 4}}}
	matrix2 := matrix{2, 2, [][]int{{5, 6}, {7, 8}}}
	matrix3 := addMat(matrix1, matrix2)
	matrix4 := multiplyMat(matrix1, matrix2)
	fmt.Println("First matrix")
	for i := 0; i < matrix3.row; i++ {
		for j := 0; j < matrix3.col; j++ {
			fmt.Print(matrix1.values[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Second matrix")

	for i := 0; i < matrix3.row; i++ {
		for j := 0; j < matrix3.col; j++ {
			fmt.Print(matrix2.values[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Added matrix")
	for i := 0; i < matrix3.row; i++ {
		for j := 0; j < matrix3.col; j++ {
			fmt.Print(matrix3.values[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Multiplied matrix")
	for i := 0; i < matrix4.row; i++ {
		for j := 0; j < matrix4.col; j++ {
			fmt.Print(matrix4.values[i][j], " ")
		}
		fmt.Println()
	}
}
