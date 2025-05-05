/*
	https://leetcode.com/problems/alternating-digit-sum/description/

	You are given a positive integer n. Each digit of n has a sign according to the following rules:

    The most significant digit is assigned a positive sign.
    Each other digit has an opposite sign to its adjacent digits.

	Return the sum of all digits with their corresponding sign.
*/

package solutions

func alternateDigitSum(n int) int {
	res := 0
	xs := []int{}

	for n != 0 {
		xs = append(xs, n%10)
		n /= 10
	}

	for i := 0; i < len(xs)>>1; i++ {
		xs[i], xs[len(xs)-i-1] = xs[len(xs)-i-1], xs[i]
	}

	for i, x := range xs {
		if i%2 != 0 {
			x = -1 * x
		}
		res += x
	}

	return res

}
