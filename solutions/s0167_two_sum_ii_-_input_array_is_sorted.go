/*
	https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

	Given a 1-indexed array of integers numbers that is already sorted in
		non-decreasing order,
	find two numbers such that they add up to a specific target number. Let these
		two numbers be numbers[index1]
	and numbers[index2] where 1 <= index1 < index2 <= numbers.length.

	Return the indices of the two numbers, index1 and index2, added by one as an
		integer array [index1, index2] of length 2.

	The tests are generated such that there is exactly one solution. You may not
		use the same element twice.
*/

package solutions

func twoSum2(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	sum := numbers[l] + numbers[r]

	for sum != target {
		if sum < target {
			l++
		} else {
			r--
		}
		sum = numbers[l] + numbers[r]
	}

	return []int{l + 1, r + 1}
}
