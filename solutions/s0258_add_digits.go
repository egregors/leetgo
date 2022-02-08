/*
	https://leetcode.com/problems/add-digits/

	Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.
*/

package solutions

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num%9 == 0 {
		return 9
	}
	return num % 9
}
