package main

import (
	"fmt"
)

func main() {
	fmt.Println(cuttingRope(120))
}
func cuttingRope(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	} else if n%3 == 0 {
		return qq(3, n/3)
	} else if n%3 == 1 {
		return qq(3, n/3-1) * 2 % 1000000007 * 2 % 1000000007
	} else if n%3 == 2 {
		return qq(3, n/3) * 2 % 1000000007
	}
	return 0
}
func qq(x, n int) int {
	r := 1
	for i := 0; i < n; i++ {
		r = r * x % 1000000007
	}
	return r
}
