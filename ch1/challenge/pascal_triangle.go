package main

import "fmt"

func main() {
	numRows := 7 // set the number of rows

	// create a two-dimensional slice to store the triangle values
	triangle := make([][]int, numRows)

	// fill in the triangle slice with values
	for i := 0; i < numRows; i++ {
		triangle[i] = make([]int, i+2)
		triangle[i][0], triangle[i][i] = 1, 1
		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}
	}

	// print the triangle
	for i := 0; i < numRows; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%d ", triangle[i][j])
		}
		fmt.Println()
	}
}
