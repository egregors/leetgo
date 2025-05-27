/*
	https://leetcode.com/problems/product-of-array-except-self/

	Given an integer array nums, return an array answer such that answer[i] is
		equal to the
	product of all the elements of nums except nums[i].

	The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit
		integer.

	You must write an algorithm that runs in O(n) time and without using the
		division operation.
*/

package solutions

func productExceptSelf(nums []int) []int {
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	l[0] = 1
	r[len(r)-1] = 1
	res := make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1] * nums[i-1]
		r[len(r)-1-i] = nums[len(nums)-i] * r[len(r)-i]
	}

	for i := 0; i < len(nums); i++ {
		res[i] = l[i] * r[i]
	}

	return res
}
