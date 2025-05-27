/*
	https://leetcode.com/problems/special-array-i/

	An array is considered special if the parity of every pair of adjacent elements
		is
	different. In other words, one element in each pair must be even, and the other
		must be odd.

	You are given an array of integers nums. Return true if nums is a special
		array, otherwise,
	return false.
*/

package solutions

func isArraySpecial(nums []int) bool {
	var even bool
	if nums[0]%2 == 0 {
		even = true
	}
	for i := 1; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			if even {
				return false
			}
			even = true
		} else {
			if !even {
				return false
			}
			even = false
		}
	}

	return true
}
