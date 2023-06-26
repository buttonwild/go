package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, target int) (ams [][]int) {
	resultRes := []int{}
	//注意此处是切片，切片和数组的区别是初始化的时候是否限制长度
	var dfs func(node *TreeNode, leftValue int)
	dfs = func(node *TreeNode, leftValue int) {
		if node == nil {
			return
		}
		leftValue -= node.Val
		resultRes = append(resultRes, node.Val)
		defer func() { resultRes = resultRes[:len(resultRes)-1] }()
		//是否需要在 defer 语句或赋值语句后面使用 () 取决于匿名函数是否需要在执行时传递参数或调用函数。如果需要传递参数或调用函数，则需要使用 () 来表示函数调用；如果匿名函数没有参数，可以省略 ()。
		if node.Left == nil && node.Right == nil && leftValue == 0 {
			ams = append(ams, append([]int(nil), resultRes...))
			//使用了两次append，append([]int(nil), resultRes...)表示将resultRes存储为新的切片，后续切片操作不影响当前操作
			return
		}
		dfs(node.Left, leftValue)
		dfs(node.Right, leftValue)
	}
	dfs(root, target)
	return
}

func main() {
	// 创建一个测试二叉树
	root := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
			Left: &TreeNode{
				Val: 11,
				Left: &TreeNode{
					Val: 7,
				},
				Right: &TreeNode{
					Val: 2,
				},
			},
		},
		Right: &TreeNode{
			Val: 8,
			Left: &TreeNode{
				Val: 13,
			},
			Right: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val: 5,
				},
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
	}

	target := 22

	// 调用 pathSum 函数
	result := pathSum(root, target)

	// 打印结果
	fmt.Println(result)
}
