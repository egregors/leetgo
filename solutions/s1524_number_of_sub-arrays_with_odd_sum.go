/*
	https://leetcode.com/problems/number-of-sub-arrays-with-odd-sum/description/

	Given an array of integers arr, return the number of subarrays with an odd sum.

	Since the answer can be very large, return it modulo 109 + 7.
*/

package solutions

func numOfSubarrays(arr []int) int {
	n := len(arr)
	MOD := 1_000_000_000 + 7

	for i := 0; i < n; i++ {
		arr[i] %= 2
	}

	dpEven, dpOdd := make([]int, n), make([]int, n)
	if arr[n-1] == 0 {
		dpEven[n-1] = 1
	} else {
		dpOdd[n-1] = 1
	}

	for i := n - 2; i >= 0; i-- {
		if arr[i] == 1 {
			dpOdd[i] = (1 + dpEven[i+1]) % MOD
			dpEven[i] = dpOdd[i+1]
		} else {
			dpEven[i] = (1 + dpEven[i+1]) % MOD
			dpOdd[i] = dpOdd[i+1]
		}
	}

	c := 0
	for _, i := range dpOdd {
		c += i
		c %= MOD
	}

	return c
}
