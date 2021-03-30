package main

import (
	"fmt"
)

func main() {
	fmt.Println(findNthDigit(10000))
}
func findNthDigit(n int) int {
	if n < 10 {
		return n
	}
	var start, digit, count int
	start = 1
	digit = 1
	count = 9
	for n > count {
		n = n - count
		digit = digit + 1
		start = start * 10
		count = digit * 9 * start
	}

	num := start + (n-1)/digit
	var s1, s2 []int
	for num != 0 {
		s1 = append(s1, num%10)
		num = num / 10
	}
	for len(s1) != 0 {
		s2 = append(s2, s1[len(s1)-1])
		s1 = s1[:len(s1)-1]
	}
	res := s2[(n-1)%digit]
	return res
}
