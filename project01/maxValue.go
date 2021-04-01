package main

import "fmt"

func main() {
	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	fmt.Println(maxValue(a))
}
func maxValue(grid [][]int) int {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if i == 0 {
				if j > 0 {
					grid[i][j] = grid[i][j-1] + grid[i][j]
				}
			} else {
				if j == 0 {
					grid[i][j] = grid[i-1][j] + grid[i][j]
				} else {
					if grid[i][j-1]+grid[i][j] > grid[i-1][j]+grid[i][j] {
						grid[i][j] = grid[i][j-1] + grid[i][j]
					} else {
						grid[i][j] = grid[i-1][j] + grid[i][j]
					}
				}
			}
		}
	}
	return grid[len(grid)-1][len(grid[0])-1]
}
