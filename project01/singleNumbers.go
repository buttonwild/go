package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 1}
	fmt.Println(singleNumbers(a))
}
func singleNumbers(nums []int) []int {
	cnums := make(map[int]int)
	for k := range nums {
		if cnums[nums[k]] == 1 {
			cnums[nums[k]] = cnums[nums[k]] + 1
		} else {
			cnums[nums[k]] = 1
		}
	}
	var a []int
	for v, z := range cnums {
		if z == 1 {
			a = append(a, v)
		}
	}
	return a
}
