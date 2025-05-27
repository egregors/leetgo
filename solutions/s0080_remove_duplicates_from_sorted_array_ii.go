/*
	https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description

		Given an integer array nums sorted in non-decreasing order, remove some
			duplicates in-place such that
	each unique element appears at most twice. The relative order of the elements
		should be kept the same.

	Since it is impossible to change the length of the array in some languages, you
		must instead have the result be
	placed in the first part of the array nums. More formally, if there are k
		elements after removing the duplicates,
	then the first k elements of nums should hold the final result. It does not
		matter what you leave beyond the first
	k elements.

	Return k after placing the final result in the first k slots of nums.

	Do not allocate extra space for another array. You must do this by modifying
		the input array in-place with O(1)
	extra memory.
*/

package solutions

func removeDuplicates80(nums []int) int {
	i, j, c := 1, 1, 1

	for i < len(nums) {
		if nums[i] == nums[i-1] {
			c++
			if c > 2 {
				i++
				continue
			}
		} else {
			c = 1
		}
		nums[j] = nums[i]
		j++
		i++
	}

	return j
}
