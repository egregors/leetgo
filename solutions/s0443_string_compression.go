/*
	https://leetcode.com/problems/string-compression/description/

	Given an array of characters chars, compress it using the following algorithm:

	Begin with an empty string s. For each group of consecutive repeating characters in chars:

		If the group's length is 1, append the character to s.
		Otherwise, append the character followed by the group's length.

	The compressed string s should not be returned separately, but instead, be stored in the input character array chars. Note that group lengths that are 10 or longer will be split into multiple characters in chars.

	After you are done modifying the input array, return the new length of the array.

	You must write an algorithm that uses only constant extra space.
*/

package solutions

import "fmt"

func compress(chars []byte) int {
	scan := 0
	write := 0
	l := len(chars)

	for scan < l {
		count := 0
		chars[write] = chars[scan]
		for scan < l && chars[write] == chars[scan] {
			count++
			scan++
		}
		if count > 1 {
			tmp := fmt.Sprintf("%d", count)
			for _, c := range []byte(tmp) {
				write++
				chars[write] = c
			}
		}
		write++
	}
	return write
}
