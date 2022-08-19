/*
	https://leetcode.com/problems/unique-morse-code-words/

	nternational Morse Code defines a standard encoding where each letter is mapped to a series of dots and
	dashes, as follows:

		'a' maps to ".-",
		'b' maps to "-...",
		'c' maps to "-.-.", and so on.

	For convenience, the full table for the 26 letters of the English alphabet is given below:

	[".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",
	".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

	Given an array of strings words where each word can be written as a concatenation of the Morse code of each letter.

		For example, "cab" can be written as "-.-..--...", which is the concatenation of "-.-.", ".-", and "-...".
	We will call such a concatenation the transformation of a word.

	Return the number of different transformations among all words we have.
*/

package solutions

func uniqueMorseRepresentations(words []string) int {
	m := make(map[int]string)
	abc := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....",
		"..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.",
		"--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-",
		"-.--", "--..",
	}
	for i, ch := range abc {
		m[i] = ch
	}

	f := func(w string) (res string) {
		for _, ch := range w {
			res += m[int(ch-'a')]
		}
		return
	}

	ans := make(map[string]struct{})
	for _, w := range words {
		ans[f(w)] = struct{}{}
	}

	return len(ans)
}
