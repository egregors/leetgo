/*
	https://leetcode.com/problems/single-element-in-a-sorted-array

	You are given a sorted array consisting of only integers where every element
		appears exactly twice,
	except for one element which appears exactly once.

	Return the single element that appears only once.

	Your solution must run in O(log n) time and O(1) space.
*/

package solutions

func singleNonDuplicate(nums []int) int {
	lo, hi := 0, len(nums)-1
	var mid int
	for lo < hi {
		mid = lo + (hi-lo)/2
		if mid%2 == 1 {
			mid--
		}
		if nums[mid] == nums[mid+1] {
			lo = mid + 2
		} else {
			hi = mid
		}
	}
	return nums[lo]
}
