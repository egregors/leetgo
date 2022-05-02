/*
	Given an integer array nums, move all the even integers at the beginning
	of the array followed by all the odd integers.

	Return any array that satisfies this condition.
*/

package solutions

func sortArrayByParity(nums []int) []int {
	i, j := 0, len(nums)-1
	for i < j {
		if nums[i]%2 > nums[j]%2 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[i]%2 == 0 {
			i++
		}
		if nums[j]%2 == 1 {
			j--
		}
	}
	return nums
}
