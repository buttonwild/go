package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	fmt.Println(search(a, 9))
}
func search(nums []int, target int) int {
	var header int = 0
	var end int = len(nums)
	var middle int = end / 2
	for header < end {
		if target < nums[header] || target > nums[end-1] {
			return -1
		} else if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			end = middle
			middle = header + (end-header)/2
		} else if target > nums[middle] {
			header = middle + 1
			middle = header + (end-header)/2
		}
	}
	return -1
}
