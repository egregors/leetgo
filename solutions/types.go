package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

// ListNode is licked list
type ListNode struct {
	Val  int
	Next *ListNode
}

func (l ListNode) String() string {
	next := "nil"
	if l.Next != nil {
		next = l.Next.String()
	}
	return fmt.Sprintf("%d -> %s", l.Val, next)
}

// ToSlice accumulate linked list values and append them into int slice
func (l *ListNode) ToSlice() []int {
	var vals []int
	curr := l
	for curr != nil {
		vals = append(vals, curr.Val)
		curr = curr.Next
	}
	return vals
}

// Node is just a node but with Next field
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// TreeNode just a node of binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode builds a TreeNode tree from LeetCode-style string representation
func NewTreeNode(data string) *TreeNode {
	if len(data) < 3 {
		return nil
	}

	data = strings.TrimSuffix(strings.TrimPrefix(data, "["), "]")
	nodes := strings.Split(data, ",")

	rootVal, _ := strconv.Atoi(nodes[0])
	root := &TreeNode{rootVal, nil, nil}

	buildTreeNode([]*TreeNode{root}, nodes[1:])
	return root
}

func buildTreeNode(q []*TreeNode, all []string) {
	if len(all) == 0 {
		return
	}
	var nextQ []*TreeNode

	for len(q) > 0 {
		n := q[0]
		q = q[1:]

		if n == nil {
			continue
		}

		l, r := all[0], all[1]
		all = all[2:]
		lVal, lErr := strconv.Atoi(l)
		rVal, rErr := strconv.Atoi(r)

		if lErr != nil {
			nextQ = append(nextQ, nil)
		} else {
			left := &TreeNode{
				lVal,
				nil,
				nil,
			}
			n.Left = left
			nextQ = append(nextQ, left)
		}

		if rErr != nil {
			nextQ = append(nextQ, nil)
		} else {
			right := &TreeNode{
				rVal,
				nil,
				nil,
			}
			n.Right = right
			nextQ = append(nextQ, right)
		}
	}

	buildTreeNode(nextQ, all)
}

func (t TreeNode) String() string {
	return fmt.Sprintf("[%s]", strings.Join(t.serialize([]*TreeNode{&t}), ","))
}

func (t *TreeNode) serialize(q []*TreeNode) []string {
	if len(q) == 0 {
		return []string{}
	}
	var nextQ []*TreeNode
	var level []string
	var isLevelValid bool

	for len(q) > 0 {
		n := q[0]
		q = q[1:]
		if n == nil {
			level = append(level, "null")
			continue
		}
		level = append(level, strconv.Itoa(n.Val))
		isLevelValid = true
		nextQ = append(nextQ, n.Left, n.Right)
	}
	if !isLevelValid {
		level = []string{}
	}
	return append(level, leveDfs(nextQ)...)
}
