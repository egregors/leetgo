/*
	https://leetcode.com/problems/construct-binary-tree-from-inorder-and-postorder-traversal

	Given two integer arrays inorder and postorder where inorder is the inorder
	traversal of a binary tree and postorder is the postorder traversal of the
	same tree, construct and return the binary tree.
*/

package solutions

// buildTree106 should call buildTree to pass LeetCode tests
func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 || len(inorder) != len(postorder) {
		return nil
	}
	length := len(postorder)
	root := &TreeNode{
		Val:   postorder[length-1],
		Left:  nil,
		Right: nil,
	}
	var i int
	for ; inorder[i] != root.Val; i++ {
	}
	root.Left = buildTree106(inorder[:i], postorder[:i])
	root.Right = buildTree106(inorder[i+1:], postorder[i:length-1])
	return root
}
