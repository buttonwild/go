package main

import "fmt"

func main() {
	head := new(Node)
	head.Val = 7
	node1 := new(Node)
	node1.Val = 13
	node2 := new(Node)
	node2.Val = 11
	node3 := new(Node)
	node3.Val = 10
	node4 := new(Node)
	node4.Val = 1
	head.Next = node1
	head.Random = nil
	node1.Next = node2
	node1.Random = head
	node2.Next = node3
	node2.Random = node4
	node3.Next = node4
	node3.Random = node2
	node4.Next = nil
	node4.Random = head
	fmt.Println(copyRandomList(head))
}

func copyRandomList(head *Node) *Node {
	var copy *Node
	p := head
	if p == nil {
		return nil
	} else {
		for p != nil { //在所有节点下面创造新节点
			t := new(Node)
			t.Val = p.Val
			t.Next = p.Next
			t.Random = nil
			p.Next = t
			p = t.Next
		}
		p = head
		for p != nil { //复制Random
			t := p.Next
			if p.Random != nil {
				t.Random = p.Random.Next
			}
			p = t.Next
		}
		p = head
		copy = p.Next
		for p != nil { //提取复制节点
			copyt := p.Next
			if copyt.Next != nil {
				p.Next = p.Next.Next
				p = copyt.Next
				copyt.Next = p.Next
			} else {
				p.Next = nil
				copyt.Next = nil
				p = nil
			}
		}
	}
	return copy
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
