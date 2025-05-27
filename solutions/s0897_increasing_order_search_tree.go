/*
	https://leetcode.com/problems/increasing-order-search-tree/

	Given the root of a binary search tree, rearrange the tree in in-order so that
		the leftmost
	node in the tree is now the root of the tree, and every node has no left child
		and only one
	right child.
*/

package solutions

func bfs897(n *TreeNode, acc *[]int) {
	if n.Left != nil {
		bfs897(n.Left, acc)
	}
	*acc = append(*acc, n.Val)
	if n.Right != nil {
		bfs897(n.Right, acc)
	}
}

func increasingBST(root *TreeNode) *TreeNode {
	var vals []int
	bfs897(root, &vals)

	res := &TreeNode{Val: vals[0]}
	curr := res
	for _, v := range vals[1:] {
		curr.Right = &TreeNode{Val: v}
		curr = curr.Right
	}

	return res
}
