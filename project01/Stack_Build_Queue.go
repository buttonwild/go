package main

import (
	"container/list"
	"fmt"
)

func main() {
	obj := Constructor()
	fmt.Println(obj)
	fmt.Println(obj.DeleteHead())
	obj.AppendTail(5)
	obj.AppendTail(2)
	fmt.Println(obj.DeleteHead())
	fmt.Println(obj.DeleteHead())
}

//CQueue 定义结构体
type CQueue struct {
	stack1, stack2 *list.List
}

//Constructor 定义结构体初始化
func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

//AppendTail 定义添加
func (CQueue *CQueue) AppendTail(value int) {
	CQueue.stack1.PushBack(value)
}

//DeleteHead 定义删除
func (CQueue *CQueue) DeleteHead() int {
	if CQueue.stack2.Len() == 0 {
		for CQueue.stack1.Len() > 0 {
			CQueue.stack2.PushBack(CQueue.stack1.Remove(CQueue.stack1.Back()))
		}
	}
	if CQueue.stack2.Len() != 0 {
		e := CQueue.stack2.Back()
		CQueue.stack2.Remove(e)
		return e.Value.(int)
	}
	return -1
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
