package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(cuttingRope(10))
}
func cuttingRope(n int) int {
	if n == 2 || n == 3 {
		return n - 1
	} else if n%3 == 0 {
		return int(math.Pow(3, float64(n/3)))
	} else if n%3 == 1 {
		return int(math.Pow(3, float64(n/3-1)) * 2 * 2)
	} else if n%3 == 2 {
		return int(math.Pow(3, float64(n/3)) * 2)
	}
	return 0
}
