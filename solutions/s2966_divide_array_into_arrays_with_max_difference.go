/*
	https://leetcode.com/problems/divide-array-into-arrays-with-max-difference

	You are given an integer array nums of size n where n is a multiple of 3 and a
		positive integer k.

	Divide the array nums into n / 3 arrays of size 3 satisfying the following
		condition:

		The difference between any two elements in one array is less than or equal to
			k.

	Return a 2D array containing the arrays. If it is impossible to satisfy the
		conditions, return an empty array. And if there are multiple answers, return
		any of them.
*/

package solutions

import "sort"

func divideArray2966(nums []int, k int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0, len(nums)/3)

	for i := 0; i < len(nums)-2; i += 3 {
		a, b, c := nums[i], nums[i+1], nums[i+2]
		if b-a > k || c-b > k || c-a > k {
			return [][]int{}
		}
		res = append(res, []int{a, b, c})
	}

	return res
}
