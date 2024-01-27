package main

import (
	"fmt"
)

type pascalTriangle struct {
	rows   int
	values [][]int
}

func (pt *pascalTriangle) init() {
	pt.values = make([][]int, pt.rows)

	// Initialize the rows
	for i := 0; i < pt.rows; i++ {
		pt.values[i] = make([]int, i+1)
	}

	// Fill the values
	for i := 0; i < pt.rows; i++ {
		for j := 0; j < i+1; j++ {
			pt.values[i][j] = pt.pascal(i, j)
		}
	}
}

func (pt pascalTriangle) print() {
	for i := 0; i < pt.rows; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", pt.values[i][j])
		}
		fmt.Println()
	}
}

func (pt pascalTriangle) pascal(row int, col int) int {
	if col == 0 || col == row {
		return 1
	} else {
		return pt.values[row-1][col-1] + pt.values[row-1][col]
	}
}

func main() {
	fmt.Print("Enter the number of rows: ")
	var rows int
	fmt.Scan(&rows)

	var pt pascalTriangle = pascalTriangle{rows: rows}
	pt.init()
	pt.print()
}
