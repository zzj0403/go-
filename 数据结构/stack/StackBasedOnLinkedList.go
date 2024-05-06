package stack

import "fmt"

type node struct {
	value interface{}
	next  *node
}

type LinkedListStack struct {
	topNode *node
}

func NewLinkedListStack() *LinkedListStack {
	return &LinkedListStack{nil}
}

// Push 压栈
func (this *LinkedListStack) Push(v interface{}) {
	this.topNode = &node{v, this.topNode}
}

// Pop 出栈
func (this *LinkedListStack) Pop() interface{} {
	if this.IsEmpty() {
		return nil
	}
	v := this.topNode.value
	this.topNode = this.topNode.next
	return v
}

// IsEmpty 是否为空
func (this *LinkedListStack) IsEmpty() bool {
	return this.topNode == nil
}

// Top 栈顶元素
func (this *LinkedListStack) Top() interface{} {
	if this.IsEmpty() {
		return nil
	}
	return this.topNode.value
}

// Flush 清空栈
func (this *LinkedListStack) Flush() {
	this.topNode = nil
}

func (this *LinkedListStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty stack")
	} else {
		cur := this.topNode
		for cur != nil {
			fmt.Println(cur.value)
			cur = cur.next
		}
	}
}
