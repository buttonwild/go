package main

import "fmt"

func main() {
	fmt.Println(hammingWeight(00000000000000000000000010000000))
}
func hammingWeight(num uint32) int {
	n := 0
	for num > 0 {
		if num%2 == 1 {
			num /= 2
			n++
		} else {
			num /= 2
		}
	}
	return n
}
