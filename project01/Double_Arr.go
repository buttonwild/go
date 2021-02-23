package main

func main() {
	matrix := [][]int{
		{-5}}
	target := 0
	findNumberIn2DArray(matrix, target)

}

func findNumberIn2DArray(matrix [][]int, target int) bool {
	//考虑三种异常
	//第一空数组
	//第二只存在一个数
	//第三不存在
	row := 0
	cloumn := 0
	if len(matrix) == 0 {
		return false
	} else {
		cloumn = len(matrix[0]) - 1
	}

	for (row < len(matrix)) && (cloumn > -1) {
		n := matrix[row][cloumn]
		if target == n {
			return true
		}
		if target < n {
			cloumn--
		}
		if target > n {
			row++
		}

	}
	return false
}
