/*
	https://leetcode.com/problems/lexicographical-numbers/

	Given an integer n, return all the numbers in the range [1, n] sorted in
		lexicographical order.

	You must write an algorithm that runs in O(n) time and uses O(1) extra space.
*/

package solutions

func lexicalOrder(n int) []int {
	res := []int{}
	cur := 1

	for i := 0; i < n; i++ {
		res = append(res, cur)
		if cur*10 <= n {
			cur *= 10
		} else {
			for cur%10 == 9 || cur >= n {
				cur /= 10
			}
			cur++
		}
	}

	return res
}
