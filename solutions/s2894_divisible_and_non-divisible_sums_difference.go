/*
	https://leetcode.com/problems/divisible-and-non-divisible-sums-difference

	You are given positive integers n and m.

	Define two integers as follows:

		num1: The sum of all integers in the range [1, n] (both inclusive) that are not divisible by m.
		num2: The sum of all integers in the range [1, n] (both inclusive) that are divisible by m.

	Return the integer num1 - num2.
*/

package solutions

func differenceOfSums(n int, m int) (s int) {
	for i := 1; i <= n; i++ {
		if i%m != 0 {
			s += i
		} else {
			s -= i
		}
	}

	return s
}
