/*
	https://leetcode.com/problems/find-peak-element/

	A peak element is an element that is strictly greater than its neighbors.

	Given an integer array nums, find a peak element, and return its index.
	If the array contains multiple peaks, return the index to any of the peaks.

	You may imagine that nums[-1] = nums[n] = -∞.

	You must write an algorithm that runs in O(log n) time.
*/

package solutions

func findPeakElement(nums []int) int {
	n := len(nums)
	lo, hi := 0, n-1
	for lo < hi {
		mid := (lo + hi) >> 1

		if nums[mid] < nums[mid+1] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}
