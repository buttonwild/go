package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{4, 5, 3, 2, 1}
	fmt.Println(validateStackSequences(a, b))
}
func validateStackSequences(pushed []int, popped []int) bool {
	var a []int
	i := len(pushed)
	for i != 0 {
		a = append(a, pushed[0])
		pushed = pushed[1:]
		for len(a) != 0 && a[len(a)-1] == popped[0] {
			popped = popped[1:]
			a = a[:len(a)-1]
		}
		i--
	}
	if len(a) == 0 {
		return true
	}
	return false
}
