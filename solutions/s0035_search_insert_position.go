/*
	https://leetcode.com/problems/search-insert-position/

	Given a sorted array of distinct integers and a target value, return the index if
	the target is found. If not, return the index where it would be if it were inserted
	in order.

	You must write an algorithm with O(log n) runtime complexity.
*/

package solutions

func searchInsert(nums []int, target int) int {
	if target < nums[0] {
		return 0
	}

	if target > nums[len(nums)-1] {
		return len(nums)
	}

	min, max := 0, len(nums)-1
	var mid int
	for min <= max {
		mid = (min + max) >> 1
		if target < nums[mid] {
			max = mid - 1
		} else if target > nums[mid] {
			min = mid + 1
		} else {
			return mid
		}
	}

	if nums[mid] < target {
		mid++
	}

	return mid
}
