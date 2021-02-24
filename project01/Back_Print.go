package main

import "fmt"

func main() {
	//尽量不要将变量名称和结构体名称一致，可能会导致编译器混淆
	ListNode1 := new(ListNode)
	ListNode1.Val = 1
	List := new(List)
	List.size = 1
	List.head = ListNode1
	List.tail = nil
	ListNode2 := new(ListNode)
	ListNode1.Next = ListNode2
	ListNode2.Val = 3
	List.AddNode(ListNode2)
	ListNode3 := new(ListNode)
	ListNode2.Next = ListNode3
	ListNode3.Val = 2
	ListNode3.Next = nil
	List.AddNode(ListNode3)
	fmt.Println(reversePrint(ListNode1))
}

func reversePrint(head *ListNode) []int {
	//递归解决办法，问题是体积太大
	if head == nil {
		return nil
	}
	if head.Next != nil {
		list := reversePrint(head.Next)
		//没有定义是什么类型，所以类型就是函数定义的类型 []int
		list = append(list, head.Val)
		//append函数用于数组添加
		return list
		//返回数组
	}
	return []int{head.Val}
	//返回最内侧数据，为数组的第一位
}

//List is a List
// 大写名称意味着可导出，就是公有，需要写注释
type List struct {
	size uint64
	head *ListNode
	tail *ListNode
}

//ListNode is a ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

//AddNode is for AddNode
func (List *List) AddNode(ListNode *ListNode) {
	List.tail = ListNode
	List.size++
}
