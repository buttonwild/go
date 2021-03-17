package main

func main() {
	head := new(TreeNode)
	head.Val = 3
	left1 := new(TreeNode)
	left1.Val = 4
	right1 := new(TreeNode)
	right1.Val = 5
	head.Left = left1
	head.Right = right1
	left2 := new(TreeNode)
	left2.Val = 1
	right2 := new(TreeNode)
	right2.Val = 2
	left1.Left = left2
	left1.Right = right2
	head2 := new(TreeNode)
	head2.Val = 4
	left3 := new(TreeNode)
	left3.Val = 1
	head2.Left = left3
	isSubStructure(head, head2)
}
func isSubStructure(A *TreeNode, B *TreeNode) bool {
	for num:=0;B
	for i:=A;i.Right!=nil;{
		if i.Val==B.Val{
			a:=i
		}
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
