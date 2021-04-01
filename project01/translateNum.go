package main

import "fmt"

func main() {
	fmt.Println(translateNum(12258))
}
func translateNum(num int) int {
	if num < 10 {
		return 1
	}
	var r int
	if num%100 <= 25 && num%100 > 9 {
		r += translateNum(num / 100)
		r += translateNum(num / 10)
	} else {
		r += translateNum(num / 10)
	}
	return r
}
