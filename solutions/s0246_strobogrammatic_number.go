/*
	https://leetcode.com/problems/strobogrammatic-number

	Given a string num which represents an integer, return true if num is a
		strobogrammatic number.

	A strobogrammatic number is a number that looks the same when rotated 180
		degrees (looked at upside down).
*/

package solutions

func isStrobogrammatic(num string) bool {
	ps := map[rune]rune{
		'0': '0',
		'1': '1',
		'6': '9',
		'8': '8',
		'9': '6',
	}

	if len(num) == 1 {
		_, ok := ps[rune(num[0])]

		return ok && num != "6" && num != "9"
	}

	for i := 0; i <= len(num)/2; i++ {
		if p, ok := ps[rune(num[i])]; !ok {
			return false
		} else if rune(num[len(num)-1-i]) != p {
			return false
		}
	}

	return true
}
