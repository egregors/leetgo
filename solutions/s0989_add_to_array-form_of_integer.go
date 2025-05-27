/*
	https://leetcode.com/problems/add-to-array-form-of-integer/

	The array-form of an integer num is an array representing its digits in left to
		right order.

		For example, for num = 1321, the array form is [1,3,2,1].

	Given num, the array-form of an integer, and an integer k, return the
		array-form of the
	integer num + k.
*/

package solutions

func addToArrayForm(num []int, k int) []int {
	var result []int
	var carry int
	for i := len(num) - 1; i >= 0 || k > 0 || carry > 0; i-- {
		var x int
		if i >= 0 {
			x = num[i]
		}
		var y int
		if k > 0 {
			y = k % 10
			k /= 10
		}
		sum := x + y + carry
		result = append(result, sum%10)
		carry = sum / 10
	}
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
