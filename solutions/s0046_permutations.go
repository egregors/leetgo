/*
	https://leetcode.com/problems/permutations/solution/

	Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

package solutions

func backtrack(output *[][]int, nums []int, first int) {
	n := len(nums)
	if first == n {
		nextNums := make([]int, len(nums))
		copy(nextNums, nums)
		*output = append(*output, nextNums)
		return
	}
	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(output, nums, first+1)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func permute(nums []int) [][]int {
	var output [][]int
	backtrack(&output, nums, 0)
	return output
}
