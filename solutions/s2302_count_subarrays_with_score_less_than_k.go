/*
	https://leetcode.com/problems/count-subarrays-with-score-less-than-k/

	The score of an array is defined as the product of its sum and its length.

		For example, the score of [1, 2, 3, 4, 5] is (1 + 2 + 3 + 4 + 5) * 5 = 75.

	Given a positive integer array nums and an integer k, return the number of non-empty subarrays of nums
	whose score is strictly less than k.


	A subarray is a contiguous sequence of elements within an array.
*/

package solutions

func countSubarrays2302(nums []int, k int64) int64 {
	n := len(nums)
	res, total := int64(0), int64(0)
	for i, j := 0, 0; j < n; j++ {
		total += int64(nums[j])
		for i <= j && total*int64(j-i+1) >= k {
			total -= int64(nums[i])
			i++
		}
		res += int64(j - i + 1)
	}
	return res
}
