/*
	https://leetcode.com/problems/using-a-robot-to-print-the-lexicographically-smallest-string/

	You are given a string s and a robot that currently holds an empty string t.
		Apply one of the following operations until s and t are both empty:

		Remove the first character of a string s and give it to the robot. The robot
			will append this character to the string t.
		Remove the last character of a string t and give it to the robot. The robot
			will write this character on paper.

	Return the lexicographically smallest string that can be written on the paper.
*/

package solutions

func robotWithString(s string) string {
	cnt := make([]int, 26)
	for _, c := range s {
		cnt[c-'a']++
	}

	stack := []rune{}
	res := []rune{}
	minCharacter := 'a'
	for _, c := range s {
		stack = append(stack, c)
		cnt[c-'a']--
		//nolint:gosec // it's ok
		for minCharacter != 'z' && cnt[minCharacter-'a'] == 0 {
			minCharacter++
		}
		for len(stack) > 0 && stack[len(stack)-1] <= minCharacter {
			res = append(res, stack[len(stack)-1])
			stack = stack[:len(stack)-1]
		}
	}

	return string(res)
}
