/*
	https://leetcode.com/problems/permutations/solution/

	Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

package solutions

func backtrack(n int, nums []int, output *[][]int, first int) {
	if first == n {
		nextNums := make([]int, len(nums))
		copy(nextNums, nums)
		*output = append(*output, nextNums)
		return
	}
	for i := first; i < n; i++ {
		nums[first], nums[i] = nums[i], nums[first]
		backtrack(n, nums, output, first+1)
		nums[first], nums[i] = nums[i], nums[first]
	}
}

func permute(nums []int) [][]int {
	var output [][]int
	backtrack(len(nums), nums, &output, 0)
	return output
}
