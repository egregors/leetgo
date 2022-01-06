/*
	https://leetcode.com/problems/permutations/solution/

	Given an array nums of distinct integers, return all the possible permutations. You can return the answer in any order.
*/

package solutions

func bt46(res *[][]int, buf []int, seen map[int]bool, nums []int) {
	if len(buf) == len(nums) {
		cp := make([]int, len(buf))
		copy(cp, buf)
		*res = append(*res, cp)
	}

	// 1 2 3
	for i := 0; i < len(nums); i++ {
		if seen[i] {
			continue
		}

		seen[i] = true
		buf = append(buf, nums[i])
		bt46(res, buf, seen, nums)
		buf = buf[:len(buf)-1]
		seen[i] = false
	}
}

func permute(nums []int) [][]int {
	var res [][]int
	bt46(&res, []int{}, make(map[int]bool), nums)
	return res
}
