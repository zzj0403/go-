package stack

import "fmt"

/*
基于数组实现的栈
*/
type ArrayStack struct {
	data []interface{}
	top  int
}

// NewArrayStack 创建一个栈
func NewArrayStack() *ArrayStack {
	return &ArrayStack{
		data: make([]interface{}, 0),
		top:  -1,
	}
}

// Push 压栈
func (this *ArrayStack) Push(v interface{}) {
	this.top++
	if this.top >= len(this.data) {
		this.data = append(this.data, v)
	} else {
		this.data[this.top] = v
	}
}

// Pop 出栈
func (this *ArrayStack) Pop() interface{} {
	if this.top == -1 {
		return nil

	}
	v := this.data[this.top]
	this.top--
	return v
}

// IsEmpty 是否为空
func (this *ArrayStack) IsEmpty() bool {
	return this.top == -1
}

// Top 栈顶元素
func (this *ArrayStack) Top() interface{} {
	if this.top == -1 {
		return nil
	}
	return this.data[this.top]
}

// Flush 清空栈
func (this *ArrayStack) Flush() {
	this.top = -1
	this.data = make([]interface{}, 0)
}

func (this *ArrayStack) Print() {
	if this.IsEmpty() {
		fmt.Println("empty statck")
	} else {
		for i := this.top; i >= 0; i-- {
			fmt.Println(this.data[i])
		}
	}
}
