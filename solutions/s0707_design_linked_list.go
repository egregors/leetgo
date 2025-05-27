/*
	https://leetcode.com/problems/design-linked-list/

	Design your implementation of the linked list. You can choose to use a singly
		or doubly
	linked list.
	A node in a singly linked list should have two attributes: val and next. val is
		the value of
	the current node, and next is a pointer/reference to the next node.
	If you want to use the doubly linked list, you will need one more attribute
		prev to indicate
	the previous node in the linked list. Assume all nodes in the linked list are
		0-indexed.

	Implement the MyLinkedList class:

		MyLinkedList() Initializes the MyLinkedList object.
		int get(int index) Get the value of the indexth node in the linked list. If
			the index
			is invalid, return -1.
		void addAtHead(int val) Add a node of value val before the first element of
			the linked
			list. After the insertion, the new node will be the first node of the linked
				list.
		void addAtTail(int val) Append a node of value val as the last element of the
			linked list.
		void addAtIndex(int index, int val) Add a node of value val before the indexth
			node in
			the linked list. If index equals the length of the linked list, the node will
				be appended to the end of the linked list. If index is greater than the
				length, the node will not be inserted.
		void deleteAtIndex(int index) Delete the indexth node in the linked list, if
			the index is valid.
*/

//nolint:revive // it's okey
package solutions

import "fmt"

type MyListNode struct {
	Val  int
	Next *MyListNode
}

type MyLinkedList struct {
	head *MyListNode
	len  int
}

// NewMyLinkedList should call Constructor to pass LeedCode tests
func NewMyLinkedList() MyLinkedList {
	return MyLinkedList{
		head: nil,
		len:  0,
	}
}

func (l MyLinkedList) String() string {
	res := fmt.Sprintf("len: %d || ", l.len)
	curr := l.head
	for curr != nil {
		res += fmt.Sprintf("%v -> ", curr.Val)
		curr = curr.Next
	}
	return res
}

func (l *MyLinkedList) Get(index int) int {
	i, curr := 0, l.head
	for curr != nil {
		if i == index {
			return curr.Val
		}
		i++
		curr = curr.Next
	}
	return -1
}

func (l *MyLinkedList) AddAtHead(val int) {
	newHead := &MyListNode{
		Val:  val,
		Next: l.head,
	}
	l.head = newHead
	l.len++
}

func (l *MyLinkedList) AddAtTail(val int) {
	if l.len == 0 {
		l.AddAtHead(val)
		return
	}

	var last *MyListNode
	curr := l.head
	for curr != nil {
		last = curr
		curr = curr.Next
	}
	last.Next = &MyListNode{
		Val:  val,
		Next: nil,
	}
	l.len++
}

func (l *MyLinkedList) AddAtIndex(index, val int) {
	if index > l.len {
		return
	}

	if index == 0 {
		l.AddAtHead(val)
		return
	}

	if index == l.len {
		l.AddAtTail(val)
		return
	}

	i := 0
	var prev, curr *MyListNode = nil, l.head
	for curr != nil {
		if i == index {
			prev.Next = &MyListNode{
				Val:  val,
				Next: curr,
			}
			l.len++
		}
		prev = curr
		curr = curr.Next
		i++
	}
}

func (l *MyLinkedList) DeleteAtIndex(index int) {
	if index > l.len {
		return
	}

	if index == 0 {
		l.head = l.head.Next
		l.len--
		return
	}

	i := 0
	var prev, curr *MyListNode = nil, l.head

	for curr != nil {
		if i == index {
			prev.Next = curr.Next
			l.len--
			return
		}
		prev = curr
		curr = curr.Next
		i++
	}
}

/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
