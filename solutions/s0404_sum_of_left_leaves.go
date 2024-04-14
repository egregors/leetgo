package solutions

func sumOfLeftLeaves(root *TreeNode) int {
	s := 0
	var dfs func(n *TreeNode, isL bool)
	dfs = func(n *TreeNode, isL bool) {
		if n == nil {
			return
		}
		if n.Left == nil && n.Right == nil {
			if isL {
				s += n.Val
			}
			return
		}
		dfs(n.Right, false)
		dfs(n.Left, true)
	}

	dfs(root, false)

	return s
}
