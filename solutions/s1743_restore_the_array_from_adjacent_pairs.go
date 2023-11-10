/*
	https://leetcode.com/problems/restore-the-array-from-adjacent-pairs/description/

	There is an integer array nums that consists of n unique elements, but you have forgotten it. However, you do
	remember every pair of adjacent elements in nums.

	You are given a 2D integer array adjacentPairs of size n - 1 where each adjacentPairs[i] = [ui, vi] indicates that
	the elements ui and vi are adjacent in nums.

	It is guaranteed that every adjacent pair of elements nums[i] and nums[i+1] will exist in adjacentPairs, either as
	[nums[i], nums[i+1]] or [nums[i+1], nums[i]]. The pairs can appear in any order.

	Return the original array nums. If there are multiple solutions, return any of them.
*/

package solutions

import "math"

func restoreArray(adjacentPairs [][]int) []int {
	g := make(map[int][]int)

	for _, p := range adjacentPairs {
		a, b := p[0], p[1]

		if _, ok := g[a]; !ok {
			g[a] = []int{}
		}
		if _, ok := g[b]; !ok {
			g[b] = []int{}
		}

		g[a] = append(g[a], b)
		g[b] = append(g[b], a)
	}

	root := 0
	for k, v := range g {
		if len(v) == 1 {
			root = k
			break
		}
	}

	res := make([]int, len(g))
	var dfs func(int, int, int)
	dfs = func(node, prev, i int) {
		res[i] = node
		for _, n := range g[node] {
			if n != prev {
				dfs(n, node, i+1)
			}
		}
	}

	dfs(root, math.MaxInt, 0)

	return res
}
