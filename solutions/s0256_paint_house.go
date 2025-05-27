/*
	https://leetcode.com/problems/paint-house/

	There is a row of n houses, where each house can be painted one of three
		colors: red,
	blue, or green. The cost of painting each house with a certain color is
		different. You
	have to paint all the houses such that no two adjacent houses have the same
		color.

	The cost of painting each house with a certain color is represented by an n x 3
		cost
	matrix costs.

		For example, costs[0][0] is the cost of painting house 0 with the color red;
	costs[1][2] is the cost of painting house 1 with color green, and so on...

	Return the minimum cost to paint all houses.
*/

package solutions

func minCost(costs [][]int) int {
	for r := len(costs) - 2; r >= 0; r-- {
		costs[r][0] += Minimum(costs[r+1][1], costs[r+1][2])
		costs[r][1] += Minimum(costs[r+1][0], costs[r+1][2])
		costs[r][2] += Minimum(costs[r+1][0], costs[r+1][1])
	}

	return Minimum(costs[0]...)
}
