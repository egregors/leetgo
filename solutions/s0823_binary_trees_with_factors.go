/*
	https://leetcode.com/problems/binary-trees-with-factors/

	Given an array of unique integers, arr, where each integer arr[i] is strictly greater than 1.

	We make a binary tree using these integers, and each number may be used for any number of times.
	Each non-leaf node's value should be equal to the product of the values of its children.

	Return the number of binary trees we can make. The answer may be too large so return the answer
	modulo 109 + 7.
*/

package solutions

import "sort"

func numFactoredBinaryTrees(arr []int) int {
	MOD := 1_000_000_007
	n := len(arr)
	sort.Ints(arr)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	index := make(map[int]int)
	for i := 0; i < n; i++ {
		index[arr[i]] = i
	}

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				r := arr[i] / arr[j]
				if v, ok := index[r]; ok {
					dp[i] = (dp[i] + dp[j]*dp[v]) % MOD
				}
			}
		}
	}
	ans := 0
	for _, x := range dp {
		ans += x
	}

	return ans % MOD
}
