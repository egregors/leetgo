/*
	https://leetcode.com/problems/find-triangular-sum-of-an-array/

	You are given a 0-indexed integer array nums, where nums[i] is a digit between
		0 and 9 (inclusive).

	The triangular sum of nums is the value of the only element present in nums
		after the following process terminates:

		Let nums comprise of n elements. If n == 1, end the process. Otherwise, create
			a new 0-indexed integer array newNums of length n - 1.
		For each index i, where 0 <= i < n - 1, assign the value of newNums[i] as
			(nums[i] + nums[i+1]) % 10, where % denotes modulo operator.
		Replace the array nums with newNums.
		Repeat the entire process starting from step 1.

	Return the triangular sum of nums.
*/

package solutions

func triangularSum(nums []int) int {
	f := func(xs ...int) []int {
		if len(xs) == 1 {
			return xs
		}
		res := make([]int, 0, len(xs)-1)
		for i := 1; i < len(xs); i++ {
			res = append(res, (xs[i]+xs[i-1])%10)
		}

		return res
	}

	acc := nums
	for len(acc) > 1 {
		acc = f(acc...)
	}

	return acc[0]
}
