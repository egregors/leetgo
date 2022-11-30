/*
   https://leetcode.com/problems/unique-number-of-occurrences/

   Given an array of integers arr, return true if the number of occurrences of each value in the array is unique, or false otherwise.
 */

package solutions

func uniqueOccurrences(arr []int) bool {
    m := make(map[int]int)
	for _, n := range arr { m[n]++ }
	m2 := make(map[int]struct{})
	for _, v := range m {
		if _, ok := m2[v]; ok { return false }
		m2[v] = struct{}{}
	}
    return true
}
