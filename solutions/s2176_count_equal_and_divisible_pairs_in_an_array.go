/*
	https://leetcode.com/problems/count-equal-and-divisible-pairs-in-an-array/description/

	Given a 0-indexed integer array nums of length n and an integer k, return the
		number of pairs (i, j)
	where 0 <= i < j < n, such that nums[i] == nums[j] and (i * j) is divisible by
		k.
*/

package solutions

func countPairs2176(nums []int, k int) int {
	c := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i*j%k == 0 {
				c++
			}
		}
	}
	return c
}
