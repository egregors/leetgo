/*
	2523. Closest Prime Numbers in Range

	Given two positive integers left and right, find the two integers num1 and num2
		such that:

    left <= num1 < num2 <= right .
    Both num1 and num2 are prime numbers.
    num2 - num1 is the minimum amongst all other pairs satisfying the above
    	conditions.

	Return the positive integer array ans = [num1, num2]. If there are multiple
		pairs satisfying
	these conditions, return the one with the smallest num1 value. If no such
		numbers exist,
	return [-1, -1].
*/

package solutions

import "math"

func closestPrimes(left int, right int) []int {
	diff := math.MaxInt
	var num1, num2 = -1, -1

	var ps []int
	for i := left; i <= right; i++ {
		if isPrime(i) {
			ps = append(ps, i)
		}
	}

	if len(ps) < 2 {
		return []int{-1, -1}
	}

	for i := 1; i < len(ps); i++ {
		if ps[i]-ps[i-1] < diff {
			num1 = ps[i-1]
			num2 = ps[i]
			diff = num2 - num1
		}
	}

	return []int{num1, num2}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	limit := int(math.Sqrt(float64(n)))
	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}
