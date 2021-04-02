package main

import "fmt"

func main() {
	a := "dvdf"
	fmt.Println(lengthOfLongestSubstring(a))
}
func lengthOfLongestSubstring(s string) int {
	byteMap := make(map[byte]bool, 26)

	var str string
	var a, e, r, i int
	for i < len(s) {
		a = i
		e = i
		for e < len(s) {
			if byteMap[s[e]] {
				str = s[a:e]
				if len(str) > r {
					r = len(str)
				}
				byteMap = make(map[byte]bool, 26)
				// for k := range byteMap {
				// 	byteMap[k] = false
				// }
				break
			} else {
				byteMap[s[e]] = true
				e++
				if e == len(s) {
					str = s[a:e]
					if len(str) > r {
						r = len(str)
					}
				}
			}
		}
		i++
	}
	return r
}
