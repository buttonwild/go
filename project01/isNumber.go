package main

import "fmt"

func main() {
	if isNumber("1 4") == false && isNumber("1 ") && isNumber(" .") == false && isNumber("3. ") && isNumber(". 1") == false && isNumber(". ") == false && isNumber("..2") == false && isNumber(".e1") == false && isNumber(" .9") && isNumber("6+1") == false && isNumber("4e+") == false && isNumber(".-4") == false && isNumber("+.4") && isNumber("-1.") && isNumber(" .45.") == false && isNumber("-1.e49046 ") {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
}
func isNumber(s string) bool {
	hasdot := false
	hasspace := false
	hasnum := false
	var n string
	for i := 0; i < len(s); {
		if s[i] == '-' || s[i] == '+' { //判断符号
			// i1 := num(s[i+1:]) + i
			// if i1 == i+1 {
			// 	return false
			// } else {
			// 	i = i1 + 1
			// }
			if hasnum == true || hasdot == true {
				return false
			}
			if len(s[i:]) > 1 {
				if s[i+1] >= '0' && s[i+1] <= '9' {
					oi := i
					i = num(s[i+2:]) + i + 2
					n = n + s[oi:i]
				} else if s[i+1] == '.' {
					i = i + 1
				} else {
					return false
				}
			} else {
				return false
			}
		} else if s[i] >= '0' && s[i] <= '9' { //判断数字
			hasnum = true
			if hasdot == true && hasspace == true {
				return false
			}
			oi := i
			i = num(s[i+1:]) + i + 1
			n = n + s[oi:i]
		} else if s[i] == '.' { //判断小数点
			if hasdot == true {
				return false
			} else if hasspace == true {
				hasspace = false
			}
			hasdot = true
			i = i + 1
		} else if s[i] == 'e' || s[i] == 'E' { //判断幂指数是否为整数
			if i == 0 || n == "" {
				return false
			}
			if len(s[i:]) > 1 {
				if s[i+1] == '-' || s[i+1] == '+' {
					oi := i
					if len(s[i+2:]) == 0 {
						return false
					}
					i = num(s[i+2:]) + i + 2
					n = n + s[oi:i]
					if i != len(s) {
						return false
					}
				} else if s[i+1] >= '0' || s[i+1] <= '9' {
					oi := i
					i = num(s[i+1:]) + i + 1
					n = n + s[oi:i]
					for i < len(s) {
						if s[i] != ' ' {
							return false
						}
						i++
					}
					if i != len(s) {
						return false
					}
				} else {
					return false
				}

			} else {
				return false
			}
		} else if s[i] == ' ' { //判断空位
			hasspace = true
			if n != "" { //空格之前不为空
				for ck := i; ck < len(s); ck++ {
					if s[ck] != ' ' {
						return false
					}
				}
				return true
			} else { //空格之前为空
				i = i + 1
			}
		} else { //都不匹配
			return false
		}
	}
	if len(n) == 0 {
		return false
	}
	return true
}
func num(s string) int {
	i := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			break
		}
	}
	return i
}
