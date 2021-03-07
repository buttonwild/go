package main

import "fmt"

func main() {
	fmt.Println(printNumbers(2))
}
func printNumbers(n int) []int {
	x := 9
	var a []int
	for w := 1; w < n; w++ {
		x = x*10 + 9
	}
	for i := 1; i <= x; i++ {
		a = append(a, i)
	}
	return a
}
