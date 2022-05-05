/*
	https://leetcode.com/problems/max-number-of-k-sum-pairs/

	You are given an integer array nums and an integer k.

	In one operation, you can pick two numbers from the array whose sum equals k and remove them from the array.

	Return the maximum number of operations you can perform on the array.
*/

package solutions

func maxOperations(nums []int, k int) int {
	m := make(map[int]int, len(nums))
	for _, v := range nums {
		m[v]++
	}

	pairs, i := 0, 0

	for i < len(nums) {
		if count, ok := m[k-nums[i]]; ok && count > 0 {
			m[k-nums[i]]--
			m[nums[i]]--

			if m[k-nums[i]] < 0 || m[nums[i]] < 0 {
				continue
			}
			pairs++
		}
		i++
	}

	return pairs
}
