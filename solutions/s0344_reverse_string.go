/*
	https://leetcode.com/problems/reverse-string/

	Write a function that reverses a string. The input string is given as an array
		of characters s.

	You must do this by modifying the input array in-place with O(1) extra memory.
*/

package solutions

func reverseString(s []byte) {
	lo, hi := 0, len(s)-1
	for lo <= hi {
		s[lo], s[hi] = s[hi], s[lo]
		lo++
		hi--
	}
}
