package main

import "fmt"

func main() {
	head := new(ListNode)
	head.Val = 1
	list2 := new(ListNode)
	list2.Val = 2
	head.Next = list2
	list3 := new(ListNode)
	list3.Val = 3
	list2.Next = list3
	list4 := new(ListNode)
	list4.Val = 4
	list4.Next = nil
	list3.Next = list4
	fmt.Println(getKthFromEnd(head, 2))
}

func getKthFromEnd(head *ListNode, k int) *ListNode {
	f := new(ListNode)
	n := 0
	i := head
	// if head.Val == ' ' {
	// 	return f
	// } else {
	// 	n = 1
	// }
	for ; i != nil; i = i.Next {
		n++
	}
	j := n - k
	f = head
	for a := 0; a < n; a++ {

		if a < j {
			f = f.Next
		}
	}
	return f
}

type ListNode struct {
	Val  int
	Next *ListNode
}
