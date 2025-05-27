/*
	https://leetcode.com/problems/count-the-number-of-powerful-integers/description/

		You are given three integers start, finish, and limit. You are also given a
			0-indexed string s representing
	a positive integer.

	A positive integer x is called powerful if it ends with s (in other words, s is
		a suffix of x) and each digit in
	x is at most limit.

	Return the total number of powerful integers in the range [start..finish].

	A string x is a suffix of a string y if and only if x is a substring of y that
		starts from some index (including 0)
	in y and extends to the index y.length - 1. For example, 25 is a suffix of 5125
		whereas 512 is not.
*/

package solutions

import (
	"fmt"
	"strings"
)

func numberOfPowerfulInt(start int64, finish int64, limit int, s string) int64 {
	low := fmt.Sprintf("%d", start)
	high := fmt.Sprintf("%d", finish)
	n := len(high)
	low = strings.Repeat("0", n-len(low)) + low // align digits
	preLen := n - len(s)                        // prefix length
	memo := make([]int64, n)
	for i := range memo {
		memo[i] = -1
	}

	var dfs func(int, bool, bool) int64
	dfs = func(i int, limitLow, limitHigh bool) int64 {
		// recursive boundary
		if i == n {
			return 1
		}
		if !limitLow && !limitHigh && memo[i] != -1 {
			return memo[i]
		}
		lo := 0
		if limitLow {
			lo = int(low[i] - '0')
		}
		hi := 9
		if limitHigh {
			hi = int(high[i] - '0')
		}

		var res int64
		if i < preLen {
			for digit := lo; digit <= min(hi, limit); digit++ {
				res += dfs(i+1, limitLow && digit == lo, limitHigh && digit == hi)
			}
		} else {
			x := int(s[i-preLen] - '0')
			if lo <= x && x <= min(hi, limit) {
				res = dfs(i+1, limitLow && x == lo, limitHigh && x == hi)
			}
		}

		if !limitLow && !limitHigh {
			memo[i] = res
		}
		return res
	}
	return dfs(0, true, true)
}
