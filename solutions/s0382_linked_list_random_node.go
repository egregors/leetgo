/*
	https://leetcode.com/problems/linked-list-random-node

	Given a singly linked list, return a random node's value from the linked list.
	Each node must have the same probability of being chosen.

	Implement the Solution383 class:

		Solution383(ListNode head) Initializes the object with the head of the
			singly-linked
	list head.
		int getRandom() Chooses a node randomly from the list and returns its value.
			All
	the nodes of the list should be equally likely to be chosen.
*/

//nolint:revive // it's ok
package solutions

import "math/rand"

type Solution383 struct {
	nodes []*ListNode
}

// NewSolution should call Constructor to pass LeetCode tests
func NewSolution(head *ListNode) Solution383 {
	var ns []*ListNode
	curr := head
	for curr != nil {
		ns = append(ns, curr)
		curr = curr.Next
	}
	return Solution383{
		nodes: ns,
	}
}

func (s *Solution383) GetRandom() int {
	n := rand.Intn(100) + 1 //nolint:gosec // meh
	return s.nodes[n%len(s.nodes)].Val
}
