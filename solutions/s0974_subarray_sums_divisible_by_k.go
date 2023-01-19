/*
	https://leetcode.com/problems/subarray-sums-divisible-by-k/description/

	Given an integer array nums and an integer k, return the number of non-empty
	subarrays that have a sum divisible by k.

	A subarray is a contiguous part of an array.
*/

package solutions

func subarraysDivByK(nums []int, k int) int {
	prefMod, res := 0, 0
	modGroups := make([]int, k)
	modGroups[0] = 1

	for _, num := range nums {
		prefMod = (prefMod + num%k + k) % k
		res += modGroups[prefMod]
		modGroups[prefMod]++
	}

	return res
}
