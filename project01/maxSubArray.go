package main

import "fmt"

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
}
func maxSubArray(nums []int) int {
	a := len(nums)
	x := 0
	for b := 0; b < a; b++ {
		if x > 0 {
			nums[b] = x + nums[b]
			x = nums[b]
		}
		if nums[b] > 0 && x <= 0 {
			x = nums[b]
		}
	}
	x=nums[0]
	for len(nums) > 0 {
		if nums[0] > x {
			x = nums[0]
		}
		nums = nums[1:]
	}
	return x
}
