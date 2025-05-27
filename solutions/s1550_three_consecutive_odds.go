/*
	https://leetcode.com/problems/three-consecutive-odds/

	Given an integer array arr, return true if there are three consecutive odd
		numbers in the array. Otherwise, return false
*/

package solutions

func threeConsecutiveOdds(arr []int) bool {
	c := 0
	for _, n := range arr {
		if n%2 != 0 {
			c++
		} else {
			c = 0
		}
		if c == 3 {
			return true
		}
	}

	return false
}
