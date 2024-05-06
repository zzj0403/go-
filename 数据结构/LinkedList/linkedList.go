package _linkedList

import "fmt"

type Node struct {
	next  *Node
	value interface{}
}

type LinkedList struct {
	head   *Node
	length uint
}

func NewNode(v interface{}) *Node {
	return &Node{nil, v}
}

func (this *Node) GetNext() *Node {
	return this.next
}

func (this *Node) GetValue() interface{} {
	return this.value
}

func NewLinkedList() *LinkedList {
	return &LinkedList{NewNode(0), 0}
}

// 在某个节点后面插入节点
func (this *LinkedList) InsertAfter(p *Node, v interface{}) bool {
	if nil == p {
		return false
	}
	newNode := NewNode(v)
	oldNext := p.next
	p.next = newNode
	newNode.next = oldNext
	this.length++
	return true
}

// 在某个节点前面插入节点
func (this *LinkedList) InsertBefore(p *Node, v interface{}) bool {
	if nil == p || p == this.head {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	newNode := NewNode(v)
	pre.next = newNode
	newNode.next = cur
	this.length++
	return true
}

// 在链表头部插入节点
func (this *LinkedList) InsertToHead(v interface{}) bool {
	return this.InsertAfter(this.head, v)
}

// 在链表尾部插入节点
func (this *LinkedList) InsertToTail(v interface{}) bool {
	cur := this.head
	for nil != cur.next {
		cur = cur.next
	}
	return this.InsertAfter(cur, v)
}

// 通过索引查找节点
func (this *LinkedList) FindByIndex(index uint) *Node {
	if index >= this.length {
		return nil
	}
	cur := this.head.next
	var i uint = 0
	for ; i < index; i++ {
		cur = cur.next
	}
	return cur
}

// 删除传入的节点
func (this *LinkedList) DeleteNode(p *Node) bool {
	if nil == p {
		return false
	}
	cur := this.head.next
	pre := this.head
	for nil != cur {
		if cur == p {
			break
		}
		pre = cur
		cur = cur.next
	}
	if nil == cur {
		return false
	}
	pre.next = p.next
	p = nil
	this.length--
	return true
}

// 打印链表
func (this *LinkedList) Print() {
	cur := this.head.next
	format := ""
	for nil != cur {
		format += fmt.Sprintf("%+v", cur.GetValue())
		cur = cur.next
		if nil != cur {
			format += "->"
		}
	}
	fmt.Println(format)
}
