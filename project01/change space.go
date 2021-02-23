package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "We are happy."
	fmt.Println(replaceSpace(a))
}
func replaceSpace(s string) string {
	//注意不能返回s
	//返回的是函数结果
	//strings.Replace(s, " ", "%20", -1)
	//return s
	return strings.Replace(s, " ", "%20", -1)
}
