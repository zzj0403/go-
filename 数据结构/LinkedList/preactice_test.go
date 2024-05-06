package _linkedList

import "testing"

var l *LinkedList

func init() {
	n6 := &Node{value: 6}
	n5 := &Node{value: 5, next: n6}
	n4 := &Node{value: 4, next: n5}
	n3 := &Node{value: 3, next: n4}
	n2 := &Node{value: 2, next: n3}
	n1 := &Node{value: 1, next: n2}
	l = &LinkedList{head: &Node{next: n1}}
}

func TestReverse(t *testing.T) {
	l.Print()
	l.Reverse()
	l.Print()
}

func TestHasCycleFloyd(t *testing.T) {
	t.Log(l.hasCycleFloyd())
	l.head.next.next.next.next = l.head.next.next.next
	t.Log(l.hasCycleFloyd())
}

func TestMergeSortedList(t *testing.T) {
	n5 := &Node{value: 9}
	n4 := &Node{value: 7, next: n5}
	n3 := &Node{value: 5, next: n4}
	n2 := &Node{value: 3, next: n3}
	n1 := &Node{value: 1, next: n2}
	l1 := &LinkedList{head: &Node{next: n1}}

	n10 := &Node{value: 10}
	n9 := &Node{value: 8, next: n10}
	n8 := &Node{value: 6, next: n9}
	n7 := &Node{value: 4, next: n8}
	n6 := &Node{value: 2, next: n7}
	l2 := &LinkedList{head: &Node{next: n6}}

	l.MergeSortedList(l1, l2).Print()
}

func TestDeleteBottomN(t *testing.T) {
	l.Print()
	l.DeleteBottomN(3)
	l.Print()
}

func TestFindMiddleNode(t *testing.T) {
	//l.DeleteBottomN(1)
	//l.DeleteBottomN(1)
	//l.DeleteBottomN(1)
	//l.DeleteBottomN(1)
	l.Print()
	t.Log(l.FindMiddleNode())
}
