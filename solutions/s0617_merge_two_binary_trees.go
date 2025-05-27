/*
	https://leetcode.com/problems/merge-two-binary-trees/

	You are given two binary trees root1 and root2.

	Imagine that when you put one of them to cover the other, some nodes of the two
		trees are overlapped
	while the others are not. You need to merge the two trees into a new binary
		tree. The merge rule is that
	if two nodes overlap, then sum node values up as the new value of the merged
		node. Otherwise, the NOT null
	node will be used as the node of the new tree.

	Return the merged tree.

	Note: The merging process must start from the root nodes of both trees.
*/

package solutions

func mergeNodesDFS(n1, n2 *TreeNode) *TreeNode {
	if n1 == nil {
		return n2
	}
	if n2 == nil {
		return n1
	}
	return &TreeNode{
		Val:   n1.Val + n2.Val,
		Left:  mergeNodesDFS(n1.Left, n2.Left),
		Right: mergeNodesDFS(n1.Right, n2.Right),
	}
}

func mergeTrees(root1, root2 *TreeNode) *TreeNode {
	return mergeNodesDFS(root1, root2)
}
