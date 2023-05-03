/*
	https://leetcode.com/problems/sign-of-the-product-of-an-array/

	There is a function signFunc(x) that returns:

		1 if x is positive.
		-1 if x is negative.
		0 if x is equal to 0.

	You are given an integer array nums. Let product be the product of all values in the array nums.

	Return signFunc(product).
*/

package solutions

func arraySign(nums []int) int {
	neg := 0
	for _, n := range nums {
		if n == 0 {
			return 0
		}
		if n < 0 {
			neg++
		}
	}

	if neg%2 == 0 {
		return 1
	}
	return -1
}
