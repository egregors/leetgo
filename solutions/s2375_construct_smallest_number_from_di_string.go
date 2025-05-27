/*
	https://leetcode.com/problems/construct-smallest-number-from-di-string/description/

	You are given a 0-indexed string pattern of length n consisting of the
		characters 'I' meaning increasing and 'D' meaning decreasing.

	A 0-indexed string num of length n + 1 is created using the following
		conditions:

		num consists of the digits '1' to '9', where each digit is used at most once.
		If pattern[i] == 'I', then num[i] < num[i + 1].
		If pattern[i] == 'D', then num[i] > num[i + 1].

	Return the lexicographically smallest possible string num that meets the
		conditions.
*/

package solutions

func smallestNumber(pattern string) string {
	stack := make([]byte, 0, len(pattern)+1)
	res := make([]byte, 0, len(pattern)+1)

	for i := 0; i <= len(pattern); i++ {
		stack = append(stack, byte('1'+i))
		if i == len(pattern) || pattern[i] == 'I' {
			for j := len(stack) - 1; j >= 0; j-- {
				res = append(res, stack[j])
			}
			stack = stack[:0]
		}
	}
	return string(res)
}
