/*
	https://leetcode.com/problems/3sum/

	Given an integer array nums, return all the triplets [nums[i], nums[j],
		nums[k]]
	such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.

	Notice that the solution set must not contain duplicate triplets.
*/

package solutions

import "sort"

func threeSum(nums []int) [][]int {
	res := &[][]int{}
	sort.Ints(nums)
	for i, n := range nums {
		if n > 0 {
			break
		}
		if i == 0 || nums[i-1] != nums[i] {
			twoSumII(nums, i, res)
		}
	}
	return *res
}

func twoSumII(nums []int, i int, res *[][]int) {
	lo, hi := i+1, len(nums)-1

	for lo < hi {
		sum := nums[i] + nums[lo] + nums[hi]
		if sum < 0 {
			lo++
		} else if sum > 0 {
			hi--
		} else {
			*res = append(*res, []int{nums[i], nums[lo], nums[hi]})
			lo++
			hi--
			for lo < hi && nums[lo] == nums[lo-1] {
				lo++
			}
		}

	}
}
