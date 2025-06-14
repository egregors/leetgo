/*
	https://leetcode.com/problems/maximum-difference-by-remapping-a-digit/

	You are given an integer num. You know that Bob will sneakily remap one of the
		10 possible digits (0 to 9) to another digit.

	Return the difference between the maximum and minimum values Bob can make by
		remapping exactly one digit in num.

	Notes:

		When Bob remaps a digit d1 to another digit d2, Bob replaces all occurrences
			of d1 in num with d2.
		Bob can remap a digit to itself, in which case num does not change.
		Bob can remap different digits for obtaining minimum and maximum values
			respectively.
		The resulting number after remapping can contain leading zeroes.
*/

package solutions

import (
	"math"
	"strconv"
	"strings"
)

func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	maxVal := 0
	minVal := math.MaxInt

	for i := 0; i < len(s); i++ {
		newS := strings.ReplaceAll(s[:], string(s[i]), "9")
		n, _ := strconv.Atoi(newS)
		maxVal = max(maxVal, n)
	}

	for i := 0; i < len(s); i++ {
		newS := strings.ReplaceAll(s[:], string(s[i]), "0")
		n, _ := strconv.Atoi(newS)
		minVal = min(minVal, n)
	}

	return maxVal - minVal
}
