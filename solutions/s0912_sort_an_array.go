/*
	https://leetcode.com/problems/sort-an-array/

	Given an array of integers nums, sort the array in ascending order.
*/

package solutions

func sortArray(nums []int) []int {
	qsort(nums)
	return nums
}

func qsort(nums []int) {
	pivot := nums[len(nums)/2]
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		for nums[lo] < pivot {
			lo++
		}
		for nums[hi] > pivot {
			hi--
		}
		if lo <= hi {
			nums[lo], nums[hi] = nums[hi], nums[lo]
			lo++
			hi--
		}
	}

	if lo < len(nums) {
		qsort(nums[:lo])
		qsort(nums[lo:])
	}
}
