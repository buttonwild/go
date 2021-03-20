package main

import "fmt"

func main() {
	head := new(TreeNode)
	head.Val = 1
	list2 := new(TreeNode)
	list2.Val = 2
	head.Left = list2
	list3 := new(TreeNode)
	list3.Val = 3
	head.Right = list3
	list4 := new(TreeNode)
	list4.Val = 4
	list2.Left = list4
	fmt.Println(levelOrder(head))
}
func levelOrder(root *TreeNode) []int {
	var a []int
	var b []TreeNode
	if root == nil {
	} else {
		if root.Left!=nil{
			b = append(b, *root.Left)
		}
		if root.Right!=nil{
			b = append(b, *root.Right)
		}
		a = append(a, root.Val)
		for len(b) != 0 {
			a = append(a, b[0].Val)

			if b[0].Left != nil {
				b = append(b, *b[0].Left)
			}
			if b[0].Right != nil {
				b = append(b, *b[0].Right)
			}
			b = b[1:]
		}
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
