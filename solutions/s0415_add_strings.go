/*
	https://leetcode.com/problems/add-strings/

	Given two non-negative integers, num1 and num2 represented as string, return the sum
	of num1 and num2 as a string.

	You must solve the problem without using any built-in library for handling large integers
	(such as BigInteger). You must also not convert the inputs to integers directly.
*/

package solutions

func addStrings(num1, num2 string) string {
	var l int
	if len(num1) > len(num2) {
		l = len(num1)
	} else {
		l = len(num2)
	}

	sum := make([]byte, l+1)
	carry := byte(0)

	for i := 0; i < l; i++ {
		a, b := byte(0), byte(0)

		if i < len(num1) {
			a = num1[len(num1)-i-1] - '0'
		}

		if i < len(num2) {
			b = num2[len(num2)-i-1] - '0'
		}

		s := a + b + carry
		carry = s / 10
		sum[len(sum)-i-1] = s%10 + '0'
	}

	if carry > 0 {
		sum[0] = '1'
	} else {
		sum = sum[1:]
	}

	return string(sum)
}
