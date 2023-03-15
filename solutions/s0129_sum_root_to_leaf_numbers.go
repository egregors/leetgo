/*
	https://leetcode.com/problems/sum-root-to-leaf-numbers

	You are given the root of a binary tree containing digits from 0 to 9 only.

	Each root-to-leaf path in the tree represents a number.

		For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.

	Return the total sum of all root-to-leaf numbers. Test cases are generated so that
	the answer will fit in a 32-bit integer.

	A leaf node is a node with no children.
*/

package solutions

func sumNumbers(root *TreeNode) int {
	var acc [][]int

	var dfs func(n *TreeNode, curr []int)
	dfs = func(n *TreeNode, curr []int) {
		if n.Left == nil && n.Right == nil {
			curr = append(curr, n.Val)
			cp := make([]int, len(curr))
			copy(cp, curr)
			acc = append(acc, cp)
			return
		}
		if n.Left != nil {
			dfs(n.Left, append(curr, n.Val))
		}
		if n.Right != nil {
			dfs(n.Right, append(curr, n.Val))
		}
	}

	dfs(root, []int{})
	var sum int
	for _, v := range acc {
		sum += arrToInt(v)
	}

	return sum
}

func arrToInt(arr []int) int {
	var res int
	for _, v := range arr {
		res = res*10 + v
	}
	return res
}
