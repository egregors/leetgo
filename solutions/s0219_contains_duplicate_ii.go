/*
	https://leetcode.com/problems/contains-duplicate-ii/

	Given an integer array nums and an integer k, return true if there are two distinct indices i and j
	in the array such that nums[i] == nums[j] and abs(i - j) <= k.
*/

package solutions

func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int)
	for i, n := range nums {
		if idx, ok := m[n]; ok && i-idx <= k {
			return true
		}
		m[n] = i
	}
	return false
}
