/*
	https://leetcode.com/problems/maximum-product-of-two-elements-in-an-array/description/

	Given the array of integers nums, you will choose two different indices i and j of that array.
	Return the maximum value of (nums[i]-1)*(nums[j]-1).
*/

package solutions

func maxProduct1464(nums []int) int {
	var n, m int

	for i := 0; i < len(nums); i++ {
		n = max(n, nums[i])
	}
	seen := false
	for i := 0; i < len(nums); i++ {
		if nums[i] == n && !seen {
			seen = true
			continue
		}

		m = max(m, nums[i])
	}

	return (n - 1) * (m - 1)
}
