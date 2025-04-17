/*
	https://leetcode.com/problems/count-the-number-of-good-subarrays/description

	Given an integer array nums and an integer k, return the number of good subarrays of nums.

	A subarray arr is good if there are at least k pairs of indices (i, j) such that i < j and arr[i] == arr[j].

	A subarray is a contiguous non-empty sequence of elements within an array.
*/

package solutions

func countGood(nums []int, k int) int64 {
	n := len(nums)
	same, r := 0, -1
	cnt := make(map[int]int)
	var ans int64
	for l := 0; l < n; l++ {
		for same < k && r+1 < n {
			r++
			same += cnt[nums[r]]
			cnt[nums[r]]++
		}
		if same >= k {
			ans += int64(n - r)
		}
		cnt[nums[l]]--
		same -= cnt[nums[l]]
	}
	return ans
}
