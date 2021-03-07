package main

import "fmt"

func main() {
	head := new(ListNode)
	head.Val = 4
	list2 := new(ListNode)
	list2.Val = 5
	head.Next = list2
	list3 := new(ListNode)
	list3.Val = 1
	list2.Next = list3
	list4 := new(ListNode)
	list4.Val = 9
	list4.Next = nil
	list3.Next = list4
	a := 5
	fmt.Println(deleteNode(head, a))
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(head *ListNode, val int) *ListNode {
	a := new(ListNode)
	b := new(ListNode)
	a = head
	b = a.Next
	for a != nil {
		if a.Val == val {
			head = head.Next
			break
		}
		if b.Val == val {
			a.Next = b.Next
			break
		}
		a = b
		b = a.Next
	}
	return head
}
