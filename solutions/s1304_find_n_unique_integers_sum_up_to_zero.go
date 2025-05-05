/*
	https://leetcode.com/problems/find-n-unique-integers-sum-up-to-zero/description/

	Given an integer n, return any array containing n unique integers such that they add up to 0.
*/

package solutions

func sumZero(n int) []int {
	res := []int{}
	if n%2 != 0 {
		res = append(res, 0)
	}
	for i := 0; i < n/2; i++ {
		res = append(res, i+1, -1*(i+1))
	}
	return res
}
