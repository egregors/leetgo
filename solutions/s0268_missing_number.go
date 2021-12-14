/*
	https://leetcode.com/problems/missing-number/

	Given an array nums containing n distinct numbers in the range [0, n], return the only number in the range that is missing from the array.
*/

package solutions

func missingNumber(nums []int) int {
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}
