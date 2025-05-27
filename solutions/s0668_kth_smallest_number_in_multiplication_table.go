/*
	https://leetcode.com/problems/kth-smallest-number-in-multiplication-table/

	Nearly everyone has used the Multiplication Table. The multiplication table of
		size m x n is an
	integer matrix mat where mat[i][j] == i * j (1-indexed).

	Given three integers m, n, and k, return the kth smallest element in the m x n
		multiplication table.
*/

package solutions

func f(val, m, n int) int {
	res := 0
	for i := 1; i <= m; i++ {
		tmp := min(val/i, n)
		res += tmp
	}
	return res
}

func findKthNumber(m, n, k int) int {
	low, high := 1, m*n+1
	for low < high {
		mid := low + (high-low)/2
		c := f(mid, m, n)
		if c >= k {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}
