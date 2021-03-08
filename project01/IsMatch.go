package main

import "fmt"

func main() {
	s := "aa"
	p := "a*"
	fmt.Println(isMatch(s, p))
}
func isMatch(s string, p string) bool {
	var b byte = 65
	var t byte = 32
	for len(s) != 0 {
		i := 0
		if s[0] == p[i] {
			s = s[1:]
			b = p[i]
			i++
			if i == len(p) {
				return false
			}
			p = p[i:]
			if p[0] != '*' {
				b = t
			}
		} else if s[0] != p[i] && p[i] == '.' {
			s = s[1:]
			b = p[i]
			i++
			if i == len(p) {
				return false
			}
			p = p[i:]
		} else if s[0] != p[i] && i < len(p)-1 && p[i+1] == '*' {
			p = p[2:]
		} else if p[i] == '*' {
			if b != s[0] && b != '.' {
				return false
			}
			s = s[1:]
			p = p[i+1:]
		}
	}
	if len(p) != 0 {
		return false
	}
	return true
}
