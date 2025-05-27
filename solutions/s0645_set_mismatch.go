/*
	https://leetcode.com/problems/set-mismatch/

	You have a set of integers s, which originally contains all the numbers from 1
		to n. Unfortunately,
	due to some error, one of the numbers in s got duplicated to another number in
		the set, which results
	in repetition of one number and loss of another number.

	You are given an integer array nums representing the data status of this set
		after the error.

	Find the number that occurs twice and the number that is missing and return
		them in the form of an array.
*/

package solutions

func findErrorNums(nums []int) []int {
	xs := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		xs[nums[i]]++
	}
	var dup, mis int
	for i := 1; i < len(xs); i++ {
		if xs[i] == 0 {
			mis = i
		}
		if xs[i] == 2 {
			dup = i
		}
	}
	return []int{dup, mis}
}
