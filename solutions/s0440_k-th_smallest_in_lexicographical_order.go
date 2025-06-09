/*
	https://leetcode.com/problems/k-th-smallest-in-lexicographical-order/

	Given two integers n and k, return the kth lexicographically smallest integer
		in the range [1, n].
*/

package solutions

func findKthNumber440(n int, k int) int {
	cur := 1
	k--

	for k > 0 {
		step := cntSteps(n, cur, cur+1)
		if step <= k {
			cur++
			k -= step
		} else {
			cur *= 10
			k--
		}
	}

	return cur
}

func cntSteps(n, pref1, pref2 int) int {
	steps := 0
	for pref1 <= n {
		steps += min(n+1, pref2) - pref1
		pref1 *= 10
		pref2 *= 10
	}
	return steps
}
