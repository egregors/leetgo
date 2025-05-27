/*
	https://leetcode.com/problems/count-of-interesting-subarrays/description

	You are given a 0-indexed integer array nums, an integer modulo, and an integer
		k.

	Your task is to find the count of subarrays that are interesting.

	A subarray nums[l..r] is interesting if the following condition holds:

		Let cnt be the number of indices i in the range [l, r] such that nums[i] %
			modulo == k. Then, cnt % modulo == k.

	Return an integer denoting the count of interesting subarrays.

	Note: A subarray is a contiguous non-empty sequence of elements within an
		array.
*/

package solutions

func countInterestingSubarrays(nums []int, modulo int, k int) int64 {
	n := len(nums)
	cnt := make(map[int]int)
	var res int64
	var prefix = 0
	cnt[0] = 1
	for i := 0; i < n; i++ {
		if nums[i]%modulo == k {
			prefix++
		}
		res += int64(cnt[(prefix-k+modulo)%modulo])
		cnt[prefix%modulo]++
	}

	return res
}
