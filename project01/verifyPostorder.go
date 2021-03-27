package main

import "fmt"

func main() {
	a := []int{1, 6, 3, 2, 5}
	fmt.Println(verifyPostorder(a))
}
func verifyPostorder(postorder []int) bool {
	var root int
	var left, right []int
	if len(postorder) >= 1 {
		root = postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		for len(postorder) != 0 && postorder[len(postorder)-1] > root {
			right = append(right, postorder[len(postorder)-1])
			postorder = postorder[:len(postorder)-1]
		}
		right = dx(right)
		for len(postorder) != 0 && postorder[len(postorder)-1] < root {
			left = append(left, postorder[len(postorder)-1])
			postorder = postorder[:len(postorder)-1]
		}
		left = dx(left)
		if len(postorder) != 0 || verifyPostorder(right) != true || verifyPostorder(left) != true {
			return false
		}
	}
	return true
}
func dx(nums []int) []int {
	var x []int
	for len(nums) != 0 {
		x = append(x, nums[len(nums)-1])
		nums = nums[:len(nums)-1]
	}
	return x
}
