package main

import "fmt"

func main() {
	s := "aaa"
	p := "ab*a*c*a"
	fmt.Println(isMatch(s, p))
}
func isMatch(s string, p string) bool {
	var f []int
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

	for len(p) != 0 {
		if p[0] != '*' {
			if len(p) >= 2 {
				if p[1] == '*' {
					p = p[2:]
				} else {
					return false
				}
			} else {
				return false
			}
		} else {
			p = p[1:]
			if len(p) != 0 {
				for z := 0; z < len(s2); z++ {
					if s2[z] == p[0] {
						f = append(f, z)
					}
				}
				for len(f) != 0 {
					if isMatch(s2[f[0]:], p) == true {
						return true
					} else {
						if len(p) >= 1 {
							f = f[1:]
						} else {
							return false
						}
					}
				}
				if len(p) >= 2 {
					if p[1] == '*' {
						p = p[2:]
					} else {
						return false
					}
				} else {
					return false
				}
			}
		}
	}
	return true
}
