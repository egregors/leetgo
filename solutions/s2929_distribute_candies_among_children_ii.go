/*
	https://leetcode.com/problems/distribute-candies-among-children-ii/

	You are given two positive integers n and limit.

	Return the total number of ways to distribute n candies among 3 children such
		that no child gets more than limit candies.
*/

package solutions

func distributeCandies(n int, limit int) int64 {
	var c int64
	m := min(n, limit)
	for i := 0; i <= m; i++ {
		c += int64(max(min(limit, n-i)-max(0, n-i-limit)+1, 0))
	}

	return c
}
