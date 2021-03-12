package main

import "fmt"

func main() {
	s := "bbbba"
	p := ".*a*a"
	fmt.Println(isMatch(s, p))
}
func isMatch(s string, p string) bool {
	var ds string //ds为全匹配的s的部分，dp为全匹配的p的部分
	var z bool    //全匹配结果
	var f []int
	ds, p, z = Match(s, p)        //结果检测，p为处理后的字段，表现为字符开头或者*开头
	s2 := s[len(ds):]             //s2表示未完全匹配的s的部分
	if len(p) == 0 && z == true { //全匹配为真
		return true
	} else if z == false {
		return false
	}
	for len(p) != 0 { //部分匹配需要继续处理
		if p[0] != '*' { //如果p不是*开头，那么就需要看是否可以化简为空，否则错误
			if len(p) >= 2 {
				if p[1] == '*' {
					p = p[2:]
				} else {
					return false
				}
			} else {
				return false
			}
		} else { //如果是*开头，那么去除掉*，检测剩余部分是否可以和s减去ds的部分全匹配
			p = p[1:]
			for len(p) != 0 {
				if s2[0] == p[0] || p[0] == '.' {
					f = append(f, 0)
				}
				if len(f) > 0 {
					z = isMatch(s2[f[0]:], p)
					if z == true {
						return true
					}
				}
				// for z := 0; z < len(p); z++ {
				// 	if s2[0] == p[z] || p[0] == '.' {
				// 		f = append(f, z)
				// 	}
				// }
				// for len(f) != 0 {
				// 	ds, p, z = Match(s2[f[0]:], p)
				// 	if z == true {
				// 		return true
				// 	} else {
				// 		if len(p) >= 1 {
				// 			f = f[1:]
				// 		} else {
				// 			return false
				// 		}
				// 	}
				// }
				if len(p) >= 2 {
					if p[1] == '*' {
						p = p[2:]
					} else {
						return false
					}
				} else {
					if p[0] == b {
						p = p[1:]
					} else {
						return false
					}
				}
			}
		}
	}
	return true
}

func Match(s string, p string) (string, string, bool) {

	var b byte = 65
	var t byte = 32
	var ds, dp string
	dp = p

	for len(s) != 0 {
		if len(p) == 0 {
			return s, dp, false
		}
		i := 0
		if s[0] == p[i] {
			if len(p) >= 2 && p[1] != '*' {
				ds = ds + string(s[0])
			}
			s = s[1:]
			b = p[i]
			i++
			p = p[i:]
			if len(p) > 0 && p[0] != '*' {
				b = t
			}
		} else if s[0] != p[i] && p[i] == '.' { //需要区别..和.*
			if len(p) >= 2 && p[1] != '*' {
				ds = ds + string(s[0])
			}
			s = s[1:]
			b = p[i]
			i++
			p = p[i:]
			if len(p) > 0 && p[0] != '*' {
				b = t
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
			return s, dp, false
		}
	}
	return ds, p, true
}
