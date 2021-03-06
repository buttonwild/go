package main

import "fmt"

func main() {
	fmt.Println(myPow(2.10000, 3))
}
func myPow(x float64, n int) float64 {
	//超时异常
	var num float64 = 1
	if x == 0 {
		return 0
	}
	if n == 1 {
		return x
	} else if n <= 0 {
		n = -n
		for n/2 > 0 || n == 1 {
			if n%2 == 1 {
				num = num * x
			}
			x = x * x
			n = n / 2
		}
		num = 1 / num
		return num
	} else if n >= 0 {
		for n/2 > 0 || n == 1 {
			if n%2 == 1 {
				num = num * x
			}
			x = x * x
			n = n / 2
		}
		return num
	}
	return 0
}
