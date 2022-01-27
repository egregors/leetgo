/*
	https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

	Given two integer arrays preorder and inorder where preorder is the preorder traversal of a
	binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
*/

package solutions

func buildTree(preorder, inorder []int) *TreeNode {
	m := make(map[int]int)
	for idx, val := range inorder {
		m[val] = idx
	}
	n := len(preorder)
	preorderIdx := 0
	return dfs105(preorder, m, &preorderIdx, 0, n-1)
}

func dfs105(preorder []int, m map[int]int, preorderIdx *int, leftInorder, rightInorder int) *TreeNode {
	if leftInorder > rightInorder {
		return nil
	}
	root := TreeNode{
		Val: preorder[*preorderIdx],
	}
	*preorderIdx++
	inorderIdx := m[root.Val]
	root.Left = dfs105(preorder, m, preorderIdx, leftInorder, inorderIdx-1)
	root.Right = dfs105(preorder, m, preorderIdx, inorderIdx+1, rightInorder)
	return &root
}
