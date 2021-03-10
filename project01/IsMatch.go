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
	s2 := s
	for len(s) != 0 {
		if len(p) == 0 {
			return false
		}
		i := 0
		if s[0] == p[i] {
			s = s[1:]
			b = p[i]
			i++
			if i != len(p) {
				p = p[i:]
				if p[0] != '*' {
					b = t
				}
			} else {
				p = p[i:]
			}

		} else if s[0] != p[i] && p[i] == '.' {
			s = s[1:]
			b = p[i]
			i++
			if i != len(p) {
				p = p[i:]
				if p[0] != '*' {
					b = t
				}
			} else {
				p = p[i:]
			}
		} else if s[0] != p[i] && i < len(p)-1 && p[i+1] == '*' {
			p = p[2:]
		} else if p[i] == '*' {
			if b != s[0] && b != '.' {
				p = p[1:]
			} else if b == s[0] || b == '.' {
				s = s[1:]
			} else if i != len(p)-1 && p[1] == b {
				p = string(p[0]) + p[i:]
			}
		} else {
			return false
		}
	}
	if len(p) != 0 && p[0] == '*' {
		p = string(b) + p
	}
	dx := t
	dx = dx & 0
	for len(p) != 0 {
		if p[len(p)-1] != s2[len(s2)-1] {
			if p[len(p)-1] == '.' {
				dx = p[len(p)-1]
				p = p[:len(p)-1]
			} else if p[len(p)-1] == '*' {
				dx = p[len(p)-1]
				p = p[:len(p)-1]
			} else if p[len(p)-1] == dx {
				p = p[:len(p)-1]
				s2 = s2[:len(s2)-1]
			} else {
				return false
			}
		} else {
			p = p[:len(p)-1]
			s2 = s2[:len(s2)-1]
		}

	}
	return true
}
