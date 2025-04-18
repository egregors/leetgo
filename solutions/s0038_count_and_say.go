/*
	https://leetcode.com/problems/count-and-say/description/

		The count-and-say sequence is a sequence of digit strings defined by the recursive formula:

		countAndSay(1) = "1"
		countAndSay(n) is the run-length encoding of countAndSay(n - 1).

	Run-length encoding (RLE) is a string compression method that works by replacing consecutive identical
	characters (repeated 2 or more times) with the concatenation of the character and the number marking the count
	of the characters (length of the run). For example, to compress the string "3322251" we replace "33" with "23",
	replace "222" with "32", replace "5" with "15" and replace "1" with "11". Thus the compressed string becomes
	"23321511".

	Given a positive integer n, return the nth element of the count-and-say sequence.
*/

package solutions

import "strconv"

func countAndSay(n int) string {
	s := "1"
	for i := 2; i <= n; i++ {
		t := ""
		c := 1
		for j := 1; j < len(s); j++ {
			if s[j] == s[j-1] {
				c++
			} else {
				t += strconv.Itoa(c) + string(s[j-1])
				c = 1
			}
		}
		t += strconv.Itoa(c) + string(s[len(s)-1])
		s = t
	}
	return s
}
