/*
	https://leetcode.com/problems/minimum-number-of-operations-to-make-elements-in-array-distinct/description/

	You are given an integer array nums. You need to ensure that the elements in
		the array are distinct. To achieve
	this, you can perform the following operation any number of times:

    Remove 3 elements from the beginning of the array. If the array has fewer
    	than 3 elements,
	remove all remaining elements.


	Note that an empty array is considered to have distinct elements. Return the
		minimum number
	of operations needed to make the elements in the array distinct.
*/

package solutions

func minimumOperations(nums []int) int {
	seen := make([]bool, 128)
	for i := len(nums) - 1; i >= 0; i-- {
		if seen[nums[i]] {
			return i/3 + 1
		}
		seen[nums[i]] = true
	}
	return 0
}
