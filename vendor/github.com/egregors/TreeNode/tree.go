package treenode

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// TreeNode binary tree node representation with Int value
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewTreeNode deserializes raw data and creates a new TreeNode
func NewTreeNode(data string) (*TreeNode, error) {
	if len(data) < 2 {
		return nil, errors.New("invalid input")
	}
	nodes := strings.Split(
		strings.TrimSuffix(strings.TrimPrefix(data, "["), "]"),
		Separator,
	)
	rootVal, err := strconv.Atoi(nodes[0])
	if err != nil {
		return nil, err
	}

	root := &TreeNode{Val: rootVal}
	err = bfsBuild(&NodeQueue{root}, nodes[1:])
	if err != nil {
		return nil, err
	}
	return root, nil
}

func (t TreeNode) String() string { return t.serialize() }

func (t TreeNode) serialize() string {
	data := bfs(&NodeQueue{&t})

	// remove redundant nulls
	i := len(data) - 1
	for data[i] == EmptyNodeMark {
		i--
	}

	return fmt.Sprintf("[%s]", strings.Join(data[:i+1], Separator))
}

// EmptyNodeMark is used to mark empty node in serialized string
const (
	EmptyNodeMark = "null"
	Separator     = ","
)

func bfs(q *NodeQueue) []string {
	var (
		nextQ        NodeQueue
		level        []string
		isEmptyLevel = true
	)

	for !q.IsEmpty() {
		n := q.Pop()
		if n != nil {
			level = append(level, strconv.Itoa(n.Val))
			nextQ.Push(n.Left, n.Right)
			isEmptyLevel = false
		} else {
			level = append(level, EmptyNodeMark)
		}
	}

	if isEmptyLevel {
		return nil
	}

	return append(level, bfs(&nextQ)...)
}

func bfsBuild(q *NodeQueue, data []string) error {
	if len(data) == 0 || q == nil {
		return nil
	}

	var nextQ NodeQueue

	for !q.IsEmpty() {
		n := q.Pop()
		if n != nil {
			// if the data tail of current level contains only nulls, they could be reduced.
			// that means, if the data becomes empty earlier than level does, there is no more nodes
			if len(data) == 0 {
				return nil
			}

			var (
				l string
				r = EmptyNodeMark
			)

			l, data = data[0], data[1:]

			if len(data) > 0 {
				r, data = data[0], data[1:]
			}

			if l != EmptyNodeMark {
				if lVal, lErr := strconv.Atoi(l); lErr == nil {
					n.Left = &TreeNode{Val: lVal}
				} else {
					return lErr
				}
			}

			if r != EmptyNodeMark {
				if rVal, rErr := strconv.Atoi(r); rErr == nil {
					n.Right = &TreeNode{Val: rVal}
				} else {
					return rErr
				}
			}

			nextQ.Push(n.Left, n.Right)
		}
	}

	return bfsBuild(&nextQ, data)
}
