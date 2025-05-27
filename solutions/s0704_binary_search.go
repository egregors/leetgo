/*
	https://leetcode.com/problems/binary-search/

	Given an array of integers nums which is sorted in ascending order, and an
		integer target,
	write a function to search target in nums. If target exists, then return its
		index.
	Otherwise, return -1.

	You must write an algorithm with O(log n) runtime complexity.
*/

package solutions

func search(nums []int, target int) int {
	min, max := 0, len(nums)-1
	for min <= max {
		midID := (min + max) / 2
		if midVal := nums[midID]; target < midVal {
			max = midID - 1
		} else if target > midVal {
			min = midID + 1
		} else {
			return midID
		}
	}
	return -1
}
