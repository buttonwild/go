package main

import "fmt"

func main() {
	fmt.Println(numWays(7))
}

func numWays(n int) int {
	//异常一超出时间限制
	//异常二求和顺序和计算顺序错误，先将数组二的值算出来，就会先运算，计算错误，不能将值为一的时候的值计算出来，一直都是值为二的
	//异常三矩阵计算的时候power2[0][0]先算出来后，算power2[0][1]会计算出错   两次
	//异常四应该是每次赋值为平方，不是每次和一相乘
	//异常五结果之间应该是相乘，不是想加
	//异常六提交的时候注意取模
	//异常七整数溢出，需要提前多次取模
	power1 := [][]int{{1, 1}, {1, 0}}
	power2 := [][]int{{1, 0}, {0, 1}}
	if n == 0 || n == 1 {
		return 1
	}
	m := n - 1
	var p00, p01, p10, p11 int
	for m/2 != 0 || m == 1 {
		if m%2 == 1 {
			p00 = power2[0][0]*power1[0][0] + power2[0][1]*power1[1][0]
			p01 = power2[0][0]*power1[0][1] + power2[0][1]*power1[1][1]
			p10 = power2[1][0]*power1[0][0] + power2[1][1]*power1[1][0]
			p11 = power2[1][0]*power1[0][1] + power2[1][1]*power1[1][1]
			power2[0][0] = p00 % 1000000007
			power2[0][1] = p01 % 1000000007
			power2[1][0] = p10 % 1000000007
			power2[1][1] = p11 % 1000000007
		}
		p00 = power1[0][0]*power1[0][0] + power1[0][1]*power1[1][0]
		p01 = power1[0][0]*power1[0][1] + power1[0][1]*power1[1][1]
		p10 = power1[1][0]*power1[0][0] + power1[1][1]*power1[1][0]
		p11 = power1[1][0]*power1[0][1] + power1[1][1]*power1[1][1]
		power1[0][0] = p00 % 1000000007
		power1[0][1] = p01 % 1000000007
		power1[1][0] = p10 % 1000000007
		power1[1][1] = p11 % 1000000007
		m /= 2
	}
	//和斐波那契数列的区别是0不一样了
	//斐波那契数列{{1}，{0}}
	//青蛙跳台阶{{1}，{1}}
	f := power2[0][0] + power2[0][1]
	return f % 1000000007
}
