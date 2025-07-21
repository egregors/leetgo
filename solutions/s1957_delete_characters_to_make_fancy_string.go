/*
	https://leetcode.com/problems/delete-characters-to-make-fancy-string/

	A fancy string is a string where no three consecutive characters are equal.

	Given a string s, delete the minimum possible number of characters from s to
		make it fancy.

	Return the final string after the deletion. It can be shown that the answer
		will always be unique.
*/

package solutions

func makeFancyString(s string) string {
	if len(s) < 3 {
		return s
	}
	res := []rune(s[:2])
	for i := 2; i < len(s); i++ {
		ch := rune(s[i])
		if ch == res[len(res)-1] && ch == res[len(res)-2] {
			continue
		}
		res = append(res, ch)
	}

	return string(res)
}
