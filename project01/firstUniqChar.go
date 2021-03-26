package main

import "fmt"

func main() {
	a := "cc"
	fmt.Println(firstUniqChar(a))
}
func firstUniqChar(s string) byte {
	var x byte
	if s != "" {
		var a1 string
		var a2 []int
		var b int
		b = len(s) - 1
		a1 = a1 + string(s[0])
		a2 = append(a2, 1)
		s = s[1:]
		for i := 0; i < b; i++ {
			for j := 0; j < len(a1); j++ {
				if s[i] == a1[j] {
					a2[j]++
					break
				} else if s[i] != a1[j] && j == len(a1)-1 {
					a1 = a1 + string(s[i])
					a2 = append(a2, 1)
					break
				}
			}
		}

		for i := 0; i < len(a2); i++ {
			if a2[i] == 1 {
				x = a1[i]
				break
			}
			if i == len(a2)-1 {
				x = ' '
			}
		}
	} else {
		x = ' '
	}

	return x
}
