/*
	https://leetcode.com/problems/count-symmetric-integers/description/

	You are given two positive integers low and high.

	An integer x consisting of 2 * n digits is symmetric if the sum of the
	first n digits of x is equal to the sum of the last n digits of x. Numbers
	with an odd number of digits are never symmetric.

	Return the number of symmetric integers in the range [low, high].
*/

package solutions

import "math"

func countSymmetricIntegers(low int, high int) int {
	c := 0
	for n := low; n <= high; n++ {
		l := toList(n)
		if len(l)%2 != 0 {
			continue
		}
		lSum, rSum := 0, 0
		for i := 0; i < len(l)/2; i++ {
			lSum += l[i]
			rSum += l[len(l)-1-i]
		}
		if lSum == rSum {
			c++
		}
	}

	return c
}

func toList(n int) []int {
	l := make([]int, 0, int(math.Floor(math.Log10(math.Abs(float64(n)))))+1)
	for n/10 > 0 {
		l = append(l, n%10)
		n /= 10
	}
	return append(l, n%10)
}
