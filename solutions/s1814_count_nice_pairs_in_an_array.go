/*
	https://leetcode.com/problems/count-nice-pairs-in-an-array/description/

	You are given an array nums that consists of non-negative integers. Let us
	define rev(x) as the reverse of the non-negative integer x.
	For example, rev(123) = 321, and rev(120) = 21. A pair of indices (i, j) is nice if it
	satisfies all of the following conditions:

    0 <= i < j < nums.length
    nums[i] + rev(nums[j]) == nums[j] + rev(nums[i])

	Return the number of nice pairs of indices. Since that number can be too
	large, return it modulo 109 + 7.
*/

package solutions

func countNicePairs(nums []int) int {
	rev := func(x int) int {
		res := 0
		for x > 0 {
			res = res*10 + x%10
			x /= 10
		}

		return res
	}

	arr := make([]int, len(nums))
	for i := range nums {
		arr[i] = nums[i] - rev(nums[i])
	}

	m := make(map[int]int)
	res := 0
	mod := int(1e9 + 7)

	for _, num := range arr {
		res = (res + m[num]) % mod
		m[num]++
	}

	return res
}
