/*
	https://leetcode.com/problems/find-unique-binary-string/description/

	Given an array of strings nums containing n unique binary strings each of length n,
	return a binary string of length n that does not appear in nums. If there are multiple
	answers, you may return any of them.
*/

package solutions

func findDifferentBinaryString(nums []string) string {
	res := make([]rune, 0, len(nums[0]))

	for i := 0; i < len(nums); i++ {
		if nums[i][i] == '0' {
			res = append(res, '1')
		} else {
			res = append(res, '0')
		}
	}

	return string(res)
}
