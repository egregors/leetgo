/*
	https://leetcode.com/problems/contains-duplicate/

	Given an integer array nums, return true if any value appears at least twice in
		the array,
	and return false if every element is distinct.
*/

package solutions

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for i := range nums {
		if m[nums[i]] {
			return true
		}
		m[nums[i]] = true
	}
	return false
}
