package main

import "fmt"

func main(){
	nums := []int{2,3,1,0,2,5,3}
	findRepeatNumber(nums)
}

func findRepeatNumber(nums []int) int {
	var fu [100000]int
	for i:=0;i<len(nums);i++{
		fu[nums[i]]++
		if fu[nums[i]] >= 2{
			fmt.Println(nums[i])
		}
	}
	return 0
}