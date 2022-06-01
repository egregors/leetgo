/*
	https://leetcode.com/problems/running-sum-of-1d-array/

	Given an array nums. We define a running sum of an array as
	runningSum[i] = sum(nums[0]…nums[i]).

	Return the running sum of nums.
*/

package solutions

func runningSum(nums []int) []int {
	res := []int{nums[0]}
	for i := 1; i < len(nums); i++ {
		res = append(res, res[i-1]+nums[i])
	}
	return res
}
