/*
	https://leetcode.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/

	A decimal number is called deci-binary if each of its digits is either 0 or 1 without any leading zeros.
	For example, 101 and 1100 are deci-binary, while 112 and 3001 are not.

	Given a string n that represents a positive decimal integer, return the minimum number of positive deci-binary
	numbers needed so that they sum up to n.
*/

package solutions

import "strconv"

func minPartitions(n string) int {
	ans := 0
	for _, ch := range n {
		d, _ := strconv.Atoi(string(ch))
		if d > ans {
			ans = d
		}
	}
	return ans
}
