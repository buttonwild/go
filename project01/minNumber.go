package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{3, 30, 34, 5, 9}
	fmt.Println(minNumber(a))
}
func minNumber(nums []int) string {
	var s string
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j])
	})
	for len(nums) != 0 {
		s = s + fmt.Sprintf("%d", nums[0])
		nums = nums[1:]
	}
	return s
}
func compare(a, b int) bool {
	str1 := fmt.Sprintf("%d%d", a, b)
	str2 := fmt.Sprintf("%d%d", b, a)
	if str1 < str2 {
		return true
	}
	return false
}
