package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	target := 5
	findNumberIn2DArray(matrix, target)

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
    row := 0
    cloumn := len(matrix[0])-1

	n := matrix[row][cloumn]
	if target==n{
        return true
    }
    else target<n {
        cloumn--
        n := matrix[row][cloumn]
    }
	return true
}
