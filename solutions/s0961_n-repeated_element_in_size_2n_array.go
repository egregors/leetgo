/*
	https://leetcode.com/problems/n-repeated-element-in-size-2n-array/

	You are given an integer array nums with the following properties:

		nums.length == 2 * n.
		nums contains n + 1 unique elements.
		Exactly one element of nums is repeated n times.

	Return the element that is repeated n times.
*/

package solutions

func repeatedNTimes(nums []int) int {
	n := len(nums) / 2
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] == n {
			return num
		}
	}
	return -1
}
