package _linkedList

/*
单链表反转
链表中环的检测
两个有序的链表合并
删除链表倒数第 n 个结点
求链表的中间结点
*/

// Reverse 单链表反转
// 学习链接https://www.bilibili.com/video/BV1uL411h7Qa/?spm_id_from=333.337.search-card.all.click&vd_source=4b27e65c46847365ef903ca46669e20a
func (this *LinkedList) Reverse() {
	if nil == this.head || nil == this.head.next || nil == this.head.next.next {
		return
	}
	var pre *Node = nil
	cur := this.head.next
	for nil != cur {
		tmp := cur.next
		cur.next = pre
		pre = cur
		cur = tmp
	}
	this.head.next = pre
}

// hasCycleFloyd 判断单链表是否有环
// https://www.bilibili.com/video/BV1Dj411b7gp/?spm_id_from=333.788&vd_source=4b27e65c46847365ef903ca46669e20a
func (this *LinkedList) hasCycleFloyd() bool {
	if nil != this.head {
		slow := this.head
		fast := this.head

		for fast != nil && fast.next != nil {
			slow = slow.next
			fast = fast.next.next

			if slow == fast {
				return true
			}
		}
	}
	return false
}

// MergeSortedList 两个有序的链表合并
// https://www.bilibili.com/video/BV1qL411X7vz/?spm_id_from=333.337.search-card.all.click&vd_source=4b27e65c46847365ef903ca46669e20a
func (this *LinkedList) MergeSortedList(l1, l2 *LinkedList) *LinkedList {
	if nil == l1.head || nil == l1.head.next {
		return l2
	}
	if nil == l2.head || nil == l2.head.next {
		return l1
	}
	l := &LinkedList{head: &Node{}}
	var pre = l.head

	for nil != l1.head.next && nil != l2.head.next {
		if l1.head.next.value.(int) <= l2.head.next.value.(int) {
			pre.next = l1.head.next
			l1.head.next = l1.head.next.next
		} else {
			pre.next = l2.head.next
			l2.head.next = l2.head.next.next
		}
		pre = pre.next
	}
	return l
}

// DeleteBottomN 删除链表中倒数第N个节点
func (this *LinkedList) DeleteBottomN(n int) {
	if n <= 0 || nil == this.head || nil == this.head.next {
		return
	}
	slow := this.head
	fast := this.head
	for i := 0; i < n-1; i++ {
		if nil == fast {
			return
		}
		fast = fast.next
	}

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	slow.next = slow.next.next

}

// FindMiddleNode 获取中间节点
func (this *LinkedList) FindMiddleNode() *Node {
	if nil == this.head || nil == this.head.next {
		return nil
	}
	slow := this.head
	fast := this.head

	for nil != fast && nil != fast.next {
		slow = slow.next
		fast = fast.next.next
	}
	if fast == nil {
		return slow
	} else {
		return slow.next
	}
}
