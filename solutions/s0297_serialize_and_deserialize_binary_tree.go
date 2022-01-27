/*
	https://leetcode.com/problems/serialize-and-deserialize-binary-tree/

	Serialization is the process of converting a data structure or object into a sequence of bits so that
	it can be stored in a file or memory buffer, or transmitted across a network connection link to be reconstructed
	later in the same or another computer environment.

	Design an algorithm to serialize and deserialize a binary tree. There is no restriction on how your
	serialization/deserialization algorithm should work. You just need to ensure that a binary tree can be
	serialized to a string and this string can be deserialized to the original tree structure.

	Clarification: The input/output format is the same as how LeetCode serializes a binary tree.
	You do not necessarily need to follow this format, so please be creative and come up with different
	approaches yourself.
*/

package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

//nolint:revive // it's ok
type Codec struct{}

// NewCodec should call Constructor to pass LeetCode tests
func NewCodec() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) serialize(root *TreeNode) string {
	return fmt.Sprintf("[%s]", strings.Join(leveDfs([]*TreeNode{root}), ","))
}

// Deserializes your encoded data to tree.
func (c *Codec) deserialize(data string) *TreeNode {
	if len(data) < 3 {
		return nil
	}
	return buildTreeFromString(data)
}

func buildTreeFromString(s string) *TreeNode {
	s = strings.TrimPrefix(s, "[")
	s = strings.TrimSuffix(s, "]")

	xs := strings.Split(s, ",")
	rootVal, _ := strconv.Atoi(xs[0])
	root := &TreeNode{rootVal, nil, nil}
	levelBfsBuild([]*TreeNode{root}, xs[1:])
	return root
}

func levelBfsBuild(q []*TreeNode, all []string) {
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

	levelBfsBuild(nextQ, all)
}

func leveDfs(q []*TreeNode) []string {
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
