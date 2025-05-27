/*
	https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/description/

		Given the root of a binary tree, return the lowest common ancestor of its
			deepest leaves.

	Recall that:

		The node of a binary tree is a leaf if and only if it has no children
		The depth of the root of the tree is 0. if the depth of a node is d, the depth
			of each of its children is d + 1.
		The lowest common ancestor of a set S of nodes, is the node A with the largest
			depth such that every node in S is in the subtree with root A.
*/

package solutions

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	_, lca := dfs1123(root)
	return lca
}

func dfs1123(root *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}

	leftDepth, leftLCA := dfs1123(root.Left)
	rightDepth, rightLCA := dfs1123(root.Right)

	if leftDepth > rightDepth {
		return leftDepth + 1, leftLCA
	}
	if leftDepth < rightDepth {
		return rightDepth + 1, rightLCA
	}
	return leftDepth + 1, root
}
