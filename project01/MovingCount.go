package main

import "fmt"

func main() {
	fmt.Println(movingCount(3, 2, 17))
}

func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	return dfs(dp, 0, 0, m, n, k)
}
func dfs(dp [][]int, g, c, m, n, k int) int {
	if g < 0 || c < 0 || g >= m || c >= n || g%10+g/10+c%10+c/10 > k || dp[g][c] == 1 {
		return 0
	}
	dp[g][c] = 1
	sum := 1
	sum += dfs(dp, g-1, c, m, n, k)
	sum += dfs(dp, g+1, c, m, n, k)
	sum += dfs(dp, g, c-1, m, n, k)
	sum += dfs(dp, g, c+1, m, n, k)
	return sum
}
