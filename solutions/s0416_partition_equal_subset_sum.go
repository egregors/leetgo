/*
	https://leetcode.com/problems/partition-equal-subset-sum/description/

	Given an integer array nums, return true if you can partition the array into two subsets such that the sum
	of the elements in both subsets is equal or false otherwise.
*/

package solutions

func canPartition(nums []int) bool {
	total := 0
	for _, n := range nums {
		total += n
	}
	if total%2 != 0 {
		return false
	}

	memo := make(map[[2]int]bool)
	subSum := total / 2
	n := len(nums)

	return dfs416(nums, subSum, n-1, memo)
}

func dfs416(nums []int, subSum, n int, memo map[[2]int]bool) bool {
	if subSum == 0 {
		return true
	}
	if n == 0 || subSum < 0 {
		return false
	}

	if res, ok := memo[[2]int{n, subSum}]; ok {
		return res
	}

	res := dfs416(nums, subSum, n-1, memo) || dfs416(nums, subSum-nums[n-1], n-1, memo)
	memo[[2]int{n, subSum}] = res

	return res
}
