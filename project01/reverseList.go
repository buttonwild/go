func reverseList(head *ListNode) *ListNode {
	var l []int
	for head != nil {
		l = append(l, head.Val)
		head = head.Next
	}
	A := new(ListNode)
	S := A
	for x := len(l) - 1; x > -1; x-- {
		if x == len(l)-1 {
			A.Val = l[x]
			A.Next = nil
		} else {
			B := new(ListNode)
			B.Val = l[x]
			B.Next = nil
			A.Next = B
			A = A.Next
		}
	}
	return S
}
