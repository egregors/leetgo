/*
	https://leetcode.com/problems/find-minimum-operations-to-make-all-elements-divisible-by-three/

	You are given an integer array nums. In one operation, you can add or subtract
		1 from any element of nums.

	Return the minimum number of operations to make all elements of nums divisible
		by 3.
*/

package solutions

func minimumOperations3190(nums []int) int {
	cnt := 0
	for _, n := range nums {
		if n%3 != 0 {
			cnt++
		}
	}
	return cnt
}
