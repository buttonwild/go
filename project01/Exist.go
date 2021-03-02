package main

import "fmt"

func main() {
	//board := [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}
	board2 := [][]byte{{'a'}, {'a'}}
	//board3 := [][]byte{{'a', 'b'}}
	//board4 := [][]byte{{'a', 'b'}, {'c', 'd'}}
	//board5 := [][]byte{{'c', 'a', 'a'}, {'a', 'a', 'a'}, {'b', 'c', 'd'}}
	//board6 := [][]byte{{'a'}}
	//word := "ABCCED"
	word2 := "aaa"
	//word3 := "ba"
	//word4 := "cdba"
	//word5 := "aab"
	//word6 := "ab"
	fmt.Println(exist(board2, word2))
}
func exist(board [][]byte, word string) bool {
	//异常一[["a"]]"ab"
	//内部多次调用，但是全局变量未重新赋值，手动赋值尝试
	//异常二[["a"]]"a"
	//边界异常
	//异常三[["a"]，["a"]]"a"
	//在异常二的基础上要多考虑，尽量让简单运算或者较优条件先执行
	//异常四[["a"]，["a"]]"aaa"
	//修改终止条件
	//异常五[['A', 'B']]"BA"
	//可重复
	c := 0
	g := 0
	for c != len(board[0]) && g != len(board) {
		n := 0
		if board[g][c] == word[n] {
			if n >= len(word) || dfs(board, g, c, word, n) {
				return true
			}
		}
		if c == len(board[0])-1 {
			g++
			c = 0
		} else {
			c++
		}
	}
	return false
}

func dfs(board [][]byte, g int, c int, word string, n int) bool {
	if n == len(word) {
		return true
	}
	if g < 0 || c < 0 || g == len(board) || c == len(board[0]) {
		return false
	}
	if board[g][c] == word[n] {
		temp := board[g][c]
		board[g][c] = ' '
		if dfs(board, g-1, c, word, n+1) ||
			dfs(board, g+1, c, word, n+1) ||
			dfs(board, g, c-1, word, n+1) ||
			dfs(board, g, c+1, word, n+1) {
			return true
		} else {
			board[g][c] = temp
		}
	}
	return false
}
