/*
	https://leetcode.com/problems/sort-colors/

	Given an array nums with n objects colored red, white, or blue, sort them
		in-place so that objects of the same
	color are adjacent, with the colors in the order red, white, and blue.

	We will use the integers 0, 1, and 2 to represent the color red, white, and
		blue, respectively.

	You must solve this problem without using the library's sort function.
*/

package solutions

import (
	"math/rand"
)

func sortColors(nums []int) {
	if len(nums) < 2 {
		return
	}

	l, r := 0, len(nums)-1

	pivot := rand.Int() % len(nums) //nolint:gosec // it's ok

	nums[pivot], nums[r] = nums[r], nums[pivot]

	for i := range nums {
		if nums[i] < nums[r] {
			nums[i], nums[l] = nums[l], nums[i]
			l++
		}
	}

	nums[l], nums[r] = nums[r], nums[l]

	sortColors(nums[:l])
	sortColors(nums[l+1:])
}
