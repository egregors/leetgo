/*
	https://leetcode.com/problems/first-unique-character-in-a-string/

	Given a string s, find the first non-repeating character in it and return its index.
	If it does not exist, return -1.
*/

package solutions

func firstUniqChar(s string) int {
	chars := make(map[rune]int)
	for _, ch := range s {
		chars[ch]++
	}
	for i, ch := range s {
		if chars[ch] == 1 {
			return i
		}
	}
	return -1
}
