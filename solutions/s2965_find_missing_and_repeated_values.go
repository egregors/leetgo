/*
	https://leetcode.com/problems/find-missing-and-repeated-values/description/

	You are given a 0-indexed 2D integer matrix grid of size n * n with values in
		the range [1, n2]. Each integer
	appears exactly once except a which appears twice and b which is missing. The
		task is to find the repeating and
	missing numbers a and b.

	Return a 0-indexed integer array ans of size 2 where ans[0] equals to a and
		ans[1] equals to b.
*/

package solutions

func findMissingAndRepeatedValues(grid [][]int) []int {
	n := len(grid)
	m := make(map[int]struct{}, n*n)
	var miss, dup int
	for ci := 0; ci < n; ci++ {
		for ri := 0; ri < n; ri++ {
			if _, ok := m[grid[ci][ri]]; ok {
				dup = grid[ci][ri]
				continue
			}
			m[grid[ci][ri]] = struct{}{}
		}
	}

	fullM := make(map[int]struct{}, n*n)
	for i := 1; i <= n*n; i++ {
		fullM[i] = struct{}{}
	}

	for k := range fullM {
		if _, ok := m[k]; !ok {
			miss = k
			break
		}
	}

	return []int{dup, miss}
}
