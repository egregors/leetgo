/*
	https://leetcode.com/problems/multiply-strings/

	Given two non-negative integers num1 and num2 represented as strings, return the product
	of num1 and num2, also represented as a string.

	Note: You must not use any built-in BigInteger library or convert the inputs to integer
	directly.
*/

package solutions

func multiply(num1, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	N := len(num1) + len(num2)
	answer := make([]int, N)

	fst := rev43(num1)
	snd := rev43(num2)

	for pl2, dg2 := range snd {
		for pl1, dg1 := range fst {
			numZeros := pl1 + pl2
			carry := answer[numZeros]

			mult := int(dg1-'0')*int(dg2-'0') + carry

			answer[numZeros] = mult % 10
			answer[numZeros+1] += mult / 10
		}
	}

	if answer[len(answer)-1] == 0 {
		answer = answer[:len(answer)-1]
	}

	res := ""
	for i := len(answer) - 1; i >= 0; i-- {
		res += string(rune(answer[i] + '0'))
	}

	return res
}

func rev43(s string) string {
	res := ""
	for i := len(s) - 1; i >= 0; i-- {
		res += string(s[i])
	}
	return res
}
