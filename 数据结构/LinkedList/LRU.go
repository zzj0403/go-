package _linkedList

import "fmt"

// 定义缓存节点结构
type CacheNode struct {
	key, value int
	prev, next *CacheNode
}

// 定义LRU缓存结构
type LRUCache struct {
	capacity   int
	size       int
	cache      map[int]*CacheNode
	head, tail *CacheNode
}

// 初始化LRU缓存
func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		size:     0,
		cache:    make(map[int]*CacheNode),
		head:     &CacheNode{},
		tail:     &CacheNode{},
	}
}

// 从LRU缓存中获取数据
func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		// 将被访问的节点移动到链表头部
		lru.moveToHead(node)
		return node.value
	}
	return -1
}

// 向LRU缓存中插入数据
func (lru *LRUCache) Put(key int, value int) {
	if node, ok := lru.cache[key]; ok {
		// 更新节点的值，并将其移动到链表头部
		node.value = value
		lru.moveToHead(node)
	} else {
		// 如果缓存容量已满，则淘汰链表尾部的数据
		if lru.size > lru.capacity {
			tail := lru.removeTail()
			delete(lru.cache, tail.key)
			lru.size--
		}
		// 创建新节点，并插入到链表头部
		newNode := &CacheNode{key: key, value: value}
		lru.cache[key] = newNode
		lru.addToHead(newNode)
		lru.size++

	}
}

// 将节点移动到链表头部
func (lru *LRUCache) moveToHead(node *CacheNode) {
	lru.removeNode(node)
	lru.addToHead(node)
}

// 从链表中移除节点
func (lru *LRUCache) removeNode(node *CacheNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

// 将节点插入到链表头部
func (lru *LRUCache) addToHead(node *CacheNode) {
	if lru.head.next == nil {
		lru.head.next = node
		node.prev = lru.head
		lru.tail.prev = node
		node.next = lru.tail
	} else {
		node.prev = lru.head
		node.next = lru.head.next
		lru.head.next.prev = node
		lru.head.next = node
	}
}

// 移除链表尾部的节点
func (lru *LRUCache) removeTail() *CacheNode {
	tail := lru.tail.prev
	lru.removeNode(tail)
	return tail
}

// 打印整个缓存列表
func (lru *LRUCache) PrintCache() {
	fmt.Print("LRU Cache: [")
	cur := lru.head.next
	for cur != nil {
		fmt.Printf("(%d,%d)", cur.key, cur.value)
		if cur.next != nil {
			fmt.Print(" -> ")
		}
		cur = cur.next
	}
	fmt.Println("]")
}
