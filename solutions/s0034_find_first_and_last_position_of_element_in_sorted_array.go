/*
	https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

	Given an array of integers nums sorted in non-decreasing order, find the
		starting and
	ending position of a given target value.

	If target is not found in the array, return [-1, -1].

	You must write an algorithm with O(log n) runtime complexity.
*/

package solutions

func findK(xs []int, k int) int {
	lo, hi := 0, len(xs)-1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if k < xs[mid] {
			hi = mid - 1
		} else if k > xs[mid] {
			lo = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func searchRange(nums []int, target int) []int {
	k := findK(nums, target)
	if k == -1 {
		return []int{-1, -1}
	}

	l, r := k, k
	for l-1 >= 0 && nums[l-1] == target {
		l--
	}
	for r+1 < len(nums) && nums[r+1] == target {
		r++
	}

	return []int{l, r}
}
