/*
	https://leetcode.com/problems/subarray-sum-equals-k/

	Given an array of integers nums and an integer k, return the total number of
	continuous subarrays whose sum equals to k.
*/

package solutions

func subarraySum(nums []int, k int) int {
	pref := make([]int, len(nums)+1)
	pref[0] = 0
	res := 0

	for i := 1; i <= len(nums); i++ {
		pref[i] = pref[i-1] + nums[i-1]
	}

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			if pref[j]-pref[i] == k {
				res++
			}
		}
	}

	return res
}
