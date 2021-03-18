package main

import "fmt"

func main() {
	// head := new(TreeNode)
	// head.Val = 1
	// left1 := new(TreeNode)
	// left1.Val = 2
	// right1 := new(TreeNode)
	// right1.Val = 3
	// head.Left = left1
	// head.Right = right1
	// left2 := new(TreeNode)
	// left2.Val = 4
	// // right2 := new(TreeNode)
	// // right2.Val = 2
	// left1.Left = left2
	// // left1.Right = right2
	// head2 := new(TreeNode)
	// head2.Val = 3
	// // left3 := new(TreeNode)
	// // left3.Val = 1
	// // head2.Left = left3
	head := new(TreeNode)
	head.Val = 4
	left1 := new(TreeNode)
	left1.Val = 2
	right1 := new(TreeNode)
	right1.Val = 3
	head.Left = left1
	head.Right = right1
	left2 := new(TreeNode)
	left2.Val = 4
	Right2 := new(TreeNode)
	Right2.Val = 5
	left1.Left = left2
	left1.Right = Right2
	Left3 := new(TreeNode)
	Left3.Val = 6
	Right3 := new(TreeNode)
	Right3.Val = 7
	right1.Left = Left3
	right1.Right = Right3
	Left4 := new(TreeNode)
	Left4.Val = 8
	Right4 := new(TreeNode)
	Right4.Val = 9
	left2.Left = Left4
	left2.Right = Right4
	head2 := new(TreeNode)
	head2.Val = 4
	left21 := new(TreeNode)
	left21.Val = 2
	right21 := new(TreeNode)
	right21.Val = 3
	head2.Left = left21
	head2.Right = right21
	fmt.Println(isSubStructure(head, head2))

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubStructure(A *TreeNode, B *TreeNode) bool {
	return A != nil && B != nil && (recur(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B))
}

func recur(A *TreeNode, B *TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil || A.Val != B.Val {
		return false
	}
	return recur(A.Left, B.Left) && recur(A.Right, B.Right)
}
