/*
	https://leetcode.com/problems/combinations/

	Given two integers n and k, return all possible combinations of k numbers out of the range [1, n].

	You may return the answer in any order.
*/

package solutions

func bt77(res *[][]int, curr []int, left, n, first int) {
	// term
	if left == 0 {
		cp := make([]int, len(curr))
		copy(cp, curr)
		*res = append(*res, cp)
		return
	}

	// base case
	for i := first; i <= n-left+1; i++ {
		bt77(res, append(curr, i), left-1, n, i+1)
	}
}

func combine(n int, k int) [][]int {
	var res [][]int
	curr := make([]int, 0, k)
	bt77(&res, curr, k, n, 1)
	return res
}
