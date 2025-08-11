/*
	https://leetcode.com/problems/range-product-queries-of-powers/

	Given a positive integer n, there exists a 0-indexed array called powers,
		composed of the minimum number of powers of 2 that sum to n. The array is
		sorted in non-decreasing order, and there is only one way to form the array.

	You are also given a 0-indexed 2D integer array queries, where queries[i] =
		[lefti, righti]. Each queries[i] represents a query where you have to find the
		product of all powers[j] with lefti <= j <= righti.

	Return an array answers, equal in length to queries, where answers[i] is the
		answer to the ith query. Since the answer to the ith query may be too large,
		each answers[i] should be returned modulo 109 + 7.
*/

package solutions

func productQueries(n int, queries [][]int) []int {
	const mod = 1000000007
	var bins []int
	rep := 1
	for n > 0 {
		if n%2 == 1 {
			bins = append(bins, rep)
		}
		n /= 2
		rep *= 2
	}

	ans := make([]int, 0, len(queries))
	for _, query := range queries {
		cur := 1
		for i := query[0]; i <= query[1]; i++ {
			cur = (cur * bins[i]) % mod
		}
		ans = append(ans, cur)
	}
	return ans
}
