/*
	https://leetcode.com/problems/removing-stars-from-a-string/

	You are given a string s, which contains stars *.

	In one operation, you can:

		Choose a star in s.
		Remove the closest non-star character to its left, as well as remove the star
			itself.

	Return the string after all stars have been removed.

	Note:

		The input will be generated such that the operation is always possible.
		It can be shown that the resulting string will always be unique.

*/

package solutions

func removeStars(s string) string {
	res := make([]rune, 0, len(s))
	for _, ch := range s {
		if ch == '*' {
			res = res[:len(res)-1]
		} else {
			res = append(res, ch)
		}
	}
	return string(res)
}
