/*
	https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

	Given a binary search tree (BST), find the lowest common ancestor (LCA) of two given nodes in the BST.

	According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes
	p and q as the lowest node in T that has both p and q as descendants (where we allow a node to be a
	descendant of itself).”
*/

package solutions

func getPath(n, target *TreeNode, path []int) []int {
	if n == nil {
		return []int{}
	}

	path = append(path, n.Val)
	if n.Val == target.Val {
		return path
	}

	if target.Val < n.Val {
		return getPath(n.Left, target, path)
	}

	if target.Val > n.Val {
		return getPath(n.Right, target, path)
	}

	return nil
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pPath, qPath := getPath(root, p, []int{}), getPath(root, q, []int{})
	for i := len(pPath) - 1; i >= 0; i-- {
		for j := len(qPath) - 1; j >= 0; j-- {
			if pPath[i] == qPath[j] {
				return &TreeNode{Val: pPath[i]}
			}
		}
	}
	return nil
}
