/*
	https://leetcode.com/problems/inorder-successor-in-bst-ii/

	Given a node in a binary search tree, return the in-order successor of that node
	in the BST. If that node has no in-order successor, return null.

	The successor of a node is the node with the smallest key greater than node.val.

	You will have direct access to the node but not to the root of the tree. Each node
	will have a reference to its parent node. Below is the definition for Node:

	class Node {
		public int val;
		public Node left;
		public Node right;
		public Node parent;
	}
*/

package solutions

import "math"

func inorderSuccessor(node *Node) *Node {
	root := node
	for root.Parent != nil {
		root = root.Parent
	}

	var nodes []*Node

	var dfs func(n *Node)
	dfs = func(n *Node) {
		if n.Left != nil {
			dfs(n.Left)
		}
		nodes = append(nodes, n)
		if n.Right != nil {
			dfs(n.Right)
		}
	}

	dfs(root)

	for i, n := range nodes {
		if n.Val == node.Val {
			min := math.MaxInt
			var res *Node
			for _, n := range nodes[i+1:] {
				if n.Val < min {
					min = n.Val
					res = n
				}
			}
			return res
		}
	}

	return nil
}
