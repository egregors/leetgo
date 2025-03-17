/*
	https://leetcode.com/problems/divide-array-into-equal-pairs/description/

	You are given an integer array nums consisting of 2 * n integers.

	You need to divide nums into n pairs such that:

		Each element belongs to exactly one pair.
		The elements present in a pair are equal.

	Return true if nums can be divided into n pairs, otherwise return false.

*/

package solutions

func divideArray(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	for _, v := range m {
		if v%2 != 0 {
			return false
		}
	}

	return true
}
