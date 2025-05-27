/*
	https://leetcode.com/problems/divide-two-integers/

	Given two integers dividend and divisor, divide two integers without using
		multiplication,
	division, and mod operator.

	The integer division should truncate toward zero, which means losing its
		fractional part.
	For example, 8.345 would be truncated to 8, and -2.7335 would be truncated to
		-2.

	Return the quotient after dividing dividend by divisor.

	Note: Assume we are dealing with an environment that could only store integers
		within the
	32-bit signed integer range: [−231, 231 − 1]. For this problem, if the
		quotient is strictly
	greater than 231 - 1, then return 231 - 1, and if the quotient is strictly less
		than -231,
	then return -231.
*/

package solutions

const (
	i32LimitP = 2147483647
	i32LimitN = -2147483648
)

func divide(dividend, divisor int) int {
	aSign, bSign := getSign(dividend), getSign(divisor)
	a, b := Abs(dividend), Abs(divisor)
	res := 0

	for a >= b {
		res++
		a -= b
	}

	if aSign != bSign {
		res *= -1
	}

	if res > 0 {
		if res > i32LimitP {
			return i32LimitP
		}
	} else {
		if res < i32LimitN {
			return i32LimitN
		}
	}

	return res
}

// getSign return true for positive numbers and false for negative
func getSign(x int) bool {
	return x >= 0
}
