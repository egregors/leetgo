/*
   https://leetcode.com/problems/plus-one/

   You are given a large integer represented as an integer array digits, where each digits[i]
   is the ith digit of the integer. The digits are ordered from most significant to least significant
   in left-to-right order. The large integer does not contain any leading 0's.

   Increment the large integer by one and return the resulting array of digits.
*/

package solutions

func plusOne(digits []int) []int {
	digits[len(digits)-1]++

	for i := len(digits) - 1; i > 0; i-- {
		if digits[i] > 9 {
			digits[i] = 0
			digits[i-1]++
		}
	}

	if digits[0] > 9 {
		digits[0] = 0
		digits = append([]int{1}, digits...)
	}

	return digits
}
