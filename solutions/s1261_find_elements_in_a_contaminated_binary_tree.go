/*
	https://leetcode.com/problems/find-elements-in-a-contaminated-binary-tree/description/

	Given a binary tree with the following rules:

		root.val == 0
		For any treeNode:
			If treeNode.val has a value x and treeNode.left != null, then treeNode.left.val == 2 * x + 1
			If treeNode.val has a value x and treeNode.right != null, then treeNode.right.val == 2 * x + 2

	Now the binary tree is contaminated, which means all treeNode.val have been changed to -1.

	Implement the FindElements class:

		FindElements(TreeNode* root) Initializes the object with a contaminated binary tree and recovers it.
		bool find(int target) Returns true if the target value exists in the recovered binary tree.

*/

package solutions

type FindElements struct {
	t *TreeNode
}

// NewFindElement must be renamed to Constructor to pass Leedcode tests
func NewFindElement(root *TreeNode) FindElements {
	root.Val = 0
	fix(root)
	return FindElements{t: root}
}

func fix(n *TreeNode) {
	if n.Left != nil {
		n.Left.Val = 2*n.Val + 1
		fix(n.Left)
	}
	if n.Right != nil {
		n.Right.Val = 2*n.Val + 2
		fix(n.Right)
	}
}

func (e *FindElements) Find(target int) bool {
	return find(e.t, target)
}

func find(n *TreeNode, t int) bool {
	if n == nil {
		return false
	}
	return n.Val == t || find(n.Left, t) || find(n.Right, t)
}
