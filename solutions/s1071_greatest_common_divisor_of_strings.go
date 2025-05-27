/*
	https://leetcode.com/problems/greatest-common-divisor-of-strings/

	For two strings s and t, we say "t divides s" if and only if s = t + ... + t
		(i.e., t is concatenated
	with itself one or more times).

	Given two strings str1 and str2, return the largest string x such that x
		divides both str1 and str2.
*/

package solutions

func gcdOfStrings(str1, str2 string) string {
	var gcd func(x, y int) int
	gcd = func(x, y int) int {
		if y == 0 {
			return x
		}
		return gcd(y, x%y)
	}

	if str1+str2 != str2+str1 {
		return ""
	}
	gcdLen := gcd(len(str1), len(str2))
	return str1[:gcdLen]
}
