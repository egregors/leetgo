/*
	https://leetcode.com/problems/find-lucky-integer-in-an-array/

	Given an array of integers arr, a lucky integer is an integer that has a
	frequency in the array equal to its value.

	Return the largest lucky integer in the array. If there is no lucky integer
	return -1.
*/

package solutions

func findLucky(arr []int) int {
	m := make(map[int]int)
	for _, n := range arr {
		m[n]++
	}

	max := -1
	for k, v := range m {
		if k == v && k > max {
			max = k
		}
	}
	return max
}
