/*
	https://leetcode.com/problems/minimum-operations-to-make-binary-array-elements-equal-to-one-i/description/

	You are given a	nums.

	You can do the following operation on the array any number of times (possibly zero):

		Choose any 3 consecutive elements from the array and flip all of them.

	Flipping an element means changing its value from 0 to 1, and from 1 to 0.

	Return the minimum number of operations required to make all elements in nums equal to 1. If it is impossible,
	return -1.
*/

package solutions

func minOperations3191(nums []int) int {
	c := 0
	for i := 2; i < len(nums); i++ {
		if nums[i-2] == 0 {
			c++
			nums[i-2] ^= 1
			nums[i-1] ^= 1
			nums[i] ^= 1
		}
	}
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum == len(nums) {
		return c
	}

	return -1
}
