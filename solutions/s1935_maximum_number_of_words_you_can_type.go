/*
	https://leetcode.com/problems/maximum-number-of-words-you-can-type/

	There is a malfunctioning keyboard where some letter keys do not work. All
		other keys on the keyboard work properly.

	Given a string text of words separated by a single space (no leading or
		trailing spaces) and a string brokenLetters of all distinct letter keys that
		are broken, return the number of words in text you can fully type using this
		keyboard.
*/

package solutions

import "strings"

func canBeTypedWords(text string, brokenLetters string) int {
	m := make(map[rune]bool)
	for _, l := range brokenLetters {
		m[l] = true
	}

	cnt := 0
LOOP:
	for _, w := range strings.Split(text, " ") {
		for _, l := range w {
			if m[l] {
				continue LOOP
			}
		}
		cnt++
	}

	return cnt
}
