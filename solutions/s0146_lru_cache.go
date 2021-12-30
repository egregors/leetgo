/*
	https://leetcode.com/problems/lru-cache/

	Design a data structure that follows the constraints of a Least Recently Used (LRU) cache.

	Implement the LRUCache class:

		LRUCache(int capacity) Initialize the LRU cache with positive size capacity.
		int get(int key) Return the value of the key if the key exists, otherwise return -1.
		void put(int key, int value) Update the value of the key if the key exists. Otherwise,
		add the key-value pair to the cache. If the number of keys exceeds the capacity from this operation,
		evict the least recently used key.

	The functions get and put must each run in O(1) average time complexity.
*/

//nolint:revive //meh
package solutions

// DLLNode is double linked list node
type DLLNode struct {
	k, v       int
	prev, next *DLLNode
}

type LRUCache struct {
	cache      map[int]*DLLNode
	cap, size  int
	head, tail *DLLNode
}

// NewLRUCache should call Constructor for leetcode tests
func NewLRUCache(capacity int) LRUCache {
	lru := LRUCache{
		cache: make(map[int]*DLLNode),
		cap:   capacity,
		size:  0,
		head:  &DLLNode{},
		tail:  &DLLNode{},
	}
	lru.head.next = lru.tail
	lru.tail.prev = lru.head
	return lru
}

func (lru *LRUCache) addNode(n *DLLNode) {
	n.prev = lru.head
	n.next = lru.head.next
	lru.head.next.prev = n
	lru.head.next = n
}

func (lru *LRUCache) removeNode(n *DLLNode) {
	prev := n.prev
	next := n.next
	prev.next = next
	next.prev = prev
}

func (lru *LRUCache) moveToHead(n *DLLNode) {
	lru.removeNode(n)
	lru.addNode(n)
}

func (lru *LRUCache) popTail() {
	res := lru.tail.prev
	lru.removeNode(res)
	delete(lru.cache, res.k)
	lru.size--
}

func (lru *LRUCache) Get(key int) int {
	if node, ok := lru.cache[key]; ok {
		lru.moveToHead(node)
		return node.v
	}
	return -1
}

func (lru *LRUCache) Put(key, value int) {
	if node, ok := lru.cache[key]; ok {
		node.v = value
		lru.moveToHead(node)
	} else {
		newNode := DLLNode{
			k: key,
			v: value,
		}
		lru.cache[key] = &newNode
		lru.addNode(&newNode)
		lru.size++

		if lru.size > lru.cap {
			lru.popTail()
		}
	}
}
