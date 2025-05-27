/*
	https://leetcode.com/problems/longest-substring-without-repeating-characters/

	Given a string s, find the length of the longest substring without repeating
		characters.
*/

package solutions

func lengthOfLongestSubstring(s string) int {
	chars := make(map[uint8]int, 128)
	left, right := 0, 0
	res := 0

	for right < len(s) {
		r := s[right]
		chars[r]++

		for chars[r] > 1 {
			l := s[left]
			chars[l]--
			left++
		}
		res = max(res, right-left+1)
		right++
	}

	return res
}
