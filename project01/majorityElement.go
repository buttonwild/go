package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 2, 2, 2, 5, 4, 2}
	fmt.Println(majorityElement(a))
}
func majorityElement(nums []int) int {
	var a [2][]int
	var b int
	b = len(nums) - 1
	a[0] = append(a[0], nums[0])
	a[1] = append(a[1], 1)
	nums = nums[1:]
	for i := 0; i < b; i++ {
		for j := 0; j < len(a[0]); j++ {
			if nums[i] == a[0][j] {
				a[1][j]++
				break
			} else if nums[i] != a[0][j] && j == len(a[0])-1 {
				a[0] = append(a[0], nums[i])
				a[1] = append(a[1], 1)
				break
			}
		}
	}
	var x, y int
	for i := 0; i < len(a[1]); i++ {
		if a[1][i] > x {
			x = a[1][i]
			y = i
		}
	}
	return a[0][y]
}
