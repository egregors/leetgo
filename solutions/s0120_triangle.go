/*
	https://leetcode.com/problems/triangle/

	Given a triangle array, return the minimum path sum from top to bottom.

	For each step, you may move to an adjacent number of the row below. More formally,
	if you are on index i on the current row, you may move to either index i or index i + 1
	on the next row.
*/

package solutions

func mins(xs ...int) int {
	m := 999999
	for _, x := range xs {
		if x < m {
			m = x
		}
	}
	return m
}

type p struct {
	r, i int
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 1 {
		return triangle[0][0]
	}

	dp := make(map[p]int)
	dp[p{0, 0}] = triangle[0][0]
	dp[p{1, 0}] = triangle[1][0] + dp[p{0, 0}]
	dp[p{1, 1}] = triangle[1][1] + dp[p{0, 0}]

	for r := 2; r < len(triangle); r++ {
		for i := 0; i < len(triangle[r]); i++ {
			var prev int
			if i == 0 {
				prev = dp[p{r - 1, 0}]
			} else if i == len(triangle[r])-1 {
				prev = dp[p{r - 1, len(triangle[r-1]) - 1}]
			} else {
				prev = min(dp[p{r - 1, i}], dp[p{r - 1, i - 1}])
			}

			dp[p{r, i}] = prev + triangle[r][i]
		}
	}

	var lastLine []int
	for k, v := range dp {
		if k.r == len(triangle)-1 {
			lastLine = append(lastLine, v)
		}
	}

	return mins(lastLine...)
}
