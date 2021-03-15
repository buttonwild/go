package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(exchange(a))
}
func exchange(nums []int) []int {
	var a []int
	var y []int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			a = append(a, 1)
			y = append(y, nums[i])
		} else {
			a = append(a, 0)
		}
	}
	for i := 0; i < len(nums); i++ {
		if a[i] == 0 {
			y = append(y, nums[i])
		}
	}
	return y
}
