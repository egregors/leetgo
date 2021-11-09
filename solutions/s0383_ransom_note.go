/*
	https://leetcode.com/problems/ransom-note/

	Given two stings ransomNote and magazine, return true if ransomNote can be
	constructed from magazine and false otherwise.

	Each letter in magazine can only be used once in ransomNote.
*/

package solutions

func canConstruct(ransomNote, magazine string) bool {
	chars := make(map[rune]int, len(magazine))
	for _, ch := range magazine {
		chars[ch]++
	}

	for _, ch := range ransomNote {
		if _, ok := chars[ch]; !ok {
			return false
		}
		chars[ch]--
		if chars[ch] < 0 {
			return false
		}
	}

	return true
}
