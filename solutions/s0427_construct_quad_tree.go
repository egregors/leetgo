/*
	https://leetcode.com/problems/construct-quad-tree/

	Given a n * n matrix grid of 0's and 1's only. We want to represent the grid
		with a Quad-Tree.

	Return the root of the Quad-Tree representing the grid.

	Notice that you can assign the value of a node to True or False when isLeaf is
		False, and both are accepted in the answer.

	A Quad-Tree is a tree data structure in which each internal node has exactly
		four children. Besides, each node has two attributes:

		val: True if the node represents a grid of 1's or False if the node represents
			a grid of 0's.
		isLeaf: True if the node is leaf node on the tree or False if the node has the
			four children.

	class Node {
		public boolean val;
		public boolean isLeaf;
		public Node topLeft;
		public Node topRight;
		public Node bottomLeft;
		public Node bottomRight;
	}

	We can construct a Quad-Tree from a two-dimensional area using the following
		steps:

		If the current grid has the same value (i.e all 1's or all 0's) set isLeaf
			True and set val to the value of the
	grid and set the four children to Null and stop.
		If the current grid has different values, set isLeaf to False and set val to
			any value and divide the current
	grid into four sub-grids as shown in the photo.
		Recurse for each of the children with the proper sub-grid.

	If you want to know more about the Quad-Tree, you can refer to the wiki.

	Quad-Tree format:

	The output represents the serialized format of a Quad-Tree using level order
		traversal, where null signifies a path
	terminator where no node exists below.

	It is very similar to the serialization of the binary tree. The only difference
		is that the node is represented as
	a list [isLeaf, val].

	If the value of isLeaf or val is True we represent it as 1 in the list [isLeaf,
		val] and if the value of isLeaf or
	val is False we represent it as 0.
*/

package solutions

//nolint:revive // it's ok
type Node427 struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node427
	TopRight    *Node427
	BottomLeft  *Node427
	BottomRight *Node427
}

func construct(grid [][]int) *Node427 {
	leaves := map[int]*Node427{
		0: {
			Val:    false,
			IsLeaf: true,
		},
		1: {
			Val:    true,
			IsLeaf: true,
		},
	}
	var dfs func(i, j, n int) *Node427
	dfs = func(i, j, n int) *Node427 {
		if n == 1 {
			return leaves[grid[i][j]]
		}
		n >>= 1
		quads := [4]*Node427{dfs(i, j, n), dfs(i, j+n, n), dfs(i+n, j, n), dfs(i+n, j+n, n)}
		for i := 0; i < len(quads); i++ {
			if !quads[i].IsLeaf || quads[i].Val != quads[0].Val {
				return &Node427{
					Val:         true, // any value
					IsLeaf:      false,
					TopLeft:     quads[0],
					TopRight:    quads[1],
					BottomLeft:  quads[2],
					BottomRight: quads[3],
				}
			}
		}
		return quads[0]
	}
	return dfs(0, 0, len(grid))
}
