/*
	https://leetcode.com/problems/binary-search-tree-iterator/

	Implement the BSTIterator class that represents an iterator over the in-order
		traversal of a binary search tree (BST):

		BSTIterator(TreeNode root) Initializes an object of the BSTIterator class. The
			root of the BST is given as part of the
		constructor. The pointer should be initialized to a non-existent number
			smaller than any element in the BST.
		boolean hasNext() Returns true if there exists a number in the traversal to
			the right of the pointer, otherwise returns false.
		int next() Moves the pointer to the right, then returns the number at the
			pointer.

	Notice that by initializing the pointer to a non-existent smallest number, the
		first call to next() will return the smallest element in the BST.

	You may assume that next() calls will always be valid. That is, there will be
		at least a next number in the in-order traversal when next() is called.
*/

//nolint:revive // it's ok
package solutions

func dfs173(n *TreeNode) []int {
	var acc []int

	if n.Left != nil {
		acc = append(acc, dfs173(n.Left)...)
	}

	acc = append(acc, n.Val)

	if n.Right != nil {
		acc = append(acc, dfs173(n.Right)...)
	}

	return acc
}

type BSTIterator struct {
	currID int
	values []int
}

// NewBSTIterator should call Constructor to pass LeetCode tests
func NewBSTIterator(root *TreeNode) BSTIterator {
	return BSTIterator{
		currID: -1,
		values: dfs173(root),
	}
}

func (i *BSTIterator) Next() int {
	i.currID++
	return i.values[i.currID]
}

func (i *BSTIterator) HasNext() bool {
	return i.currID < len(i.values)-1
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
