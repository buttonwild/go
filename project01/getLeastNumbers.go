package main

import "fmt"

func main() {
	a := []int{0, 0, 1, 2, 4, 2, 2, 3, 1, 4}
	fmt.Println(getLeastNumbers(a, 8))
}
func getLeastNumbers(arr []int, k int) []int {
	arr = BubbleSort(arr)
	arr = arr[:k]
	return arr
}
func BubbleSort(nums []int) []int {
	for a := 0; a < len(nums); a++ {
		for b := 0; b < len(nums)-a-1; b++ {
			if nums[b] > nums[b+1] {
				x := nums[b]
				nums[b] = nums[b+1]
				nums[b+1] = x
			}
		}
	}
	return nums
}
