/*
	https://leetcode.com/problems/find-a-corresponding-node-of-a-binary-tree-in-a-clone-of-that-tree/

	Given two binary trees original and cloned and given a reference to a node target in the original tree.

	The cloned tree is a copy of the original tree.

	Return a reference to the same node in the cloned tree.

	Note that you are not allowed to change any of the two trees or the target node and the answer must be a
	reference to a node in the cloned tree.
*/

package solutions

func findTargetNode(n, target *TreeNode, pred func(a, b *TreeNode) bool) *TreeNode {
	if pred(n, target) {
		return n
	}

	if n.Left != nil {
		if res := findTargetNode(n.Left, target, pred); res != nil {
			return res
		}
	}
	if n.Right != nil {
		if res := findTargetNode(n.Right, target, pred); res != nil {
			return res
		}
	}

	return nil
}

func isNodeEqual(n1, n2 *TreeNode) bool {
	if n1 == nil && n2 == nil {
		return true
	}

	if n1.Val == n2.Val {
		return isNodeEqual(n1.Left, n2.Left) && isNodeEqual(n1.Right, n2.Right)
	}

	return false
}

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	targetNode := findTargetNode(
		original,
		target,
		func(a, b *TreeNode) bool { return a == b },
	)

	return findTargetNode(
		cloned,
		targetNode,
		func(a, b *TreeNode) bool { return a.Val == b.Val && isNodeEqual(a, b) },
	)
}
