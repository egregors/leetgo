/*
	https://leetcode.com/problems/greatest-sum-divisible-by-three/

	Given an integer array nums, return the maximum possible sum of elements of the
		array such that it is divisible by three.
*/

package solutions

import "math"

func maxSumDivThree(nums []int) int {
	f := []int{0, math.MinInt, math.MinInt}
	for _, num := range nums {
		g := make([]int, 3)
		for i := 0; i < 3; i++ {
			g[(i+num)%3] = max(f[(i+num)%3], f[i]+num)
		}
		f = g
	}

	return f[0]
}
