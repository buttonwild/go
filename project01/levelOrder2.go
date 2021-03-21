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
	fmt.Println(levelOrder2(head))
}
func levelOrder2(root *TreeNode) [][]int {
	var a [][]int
	var c []int
	var d []TreeNode
	if root == nil {
	} else {
		c = append(c, root.Val)
		a = append(a, c)
		if root.Left != nil {
			d = append(d, *root.Left)
		}
		if root.Right != nil {
			d = append(d, *root.Right)
		}
		for len(d) != 0 {
			c, d = ft(d)
			a = append(a, c)
		}
	}
	return a
}

func ft(d []TreeNode) ([]int, []TreeNode) {
	var c []int
	var e []TreeNode
	for len(d) != 0 {
		c = append(c, d[0].Val)
		if d[0].Left != nil {
			e = append(e, *d[0].Left)
		}
		if d[0].Right != nil {
			e = append(e, *d[0].Right)
		}
		d = d[1:]
	}
	return c, e
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
