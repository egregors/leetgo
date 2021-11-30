/*
	https://leetcode.com/problems/majority-element/

	Given an array nums of size n, return the majority element.

	The majority element is the element that appears more than ⌊n / 2⌋ times.
	You may assume that the majority element always exists in the array.
*/

package solutions

func majorityElement(nums []int) int {
	n := len(nums) / 2
	xs := make(map[int]int)
	for _, i := range nums {
		xs[i]++
		if xs[i] > n {
			return i
		}
	}
	return -1
}
