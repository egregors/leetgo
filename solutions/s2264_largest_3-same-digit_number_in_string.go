/*
	https://leetcode.com/problems/largest-3-same-digit-number-in-string

	You are given a string num representing a large integer. An integer is good if
		it meets the following conditions:

		It is a substring of num with length 3.
		It consists of only one unique digit.

	Return the maximum good integer as a string or an empty string "" if no such
		integer exists.

	Note:

		A substring is a contiguous sequence of characters within a string.
		There may be leading zeroes in num or a good integer.
*/

package solutions

import "strings"

func largestGoodInteger(num string) string {
	patterns := []string{"999", "888", "777", "666", "555", "444", "333", "222", "111", "000"}
	for _, p := range patterns {
		if strings.Contains(num, p) {
			return p
		}
	}
	return ""
}
