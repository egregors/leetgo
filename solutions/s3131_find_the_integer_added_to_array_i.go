/*
	https://leetcode.com/problems/find-the-integer-added-to-array-i/description/

	You are given two arrays of equal length, nums1 and nums2.

	Each element in nums1 has been increased (or decreased in the case of negative) 
	by an integer, represented by the variable x.

	As a result, nums1 becomes equal to nums2. Two arrays are considered equal when 
	they contain the same integers with the same frequencies.

	Return the integer x.
*/

package solutions

import "math"

func addedInteger(nums1 []int, nums2 []int) int {
    min1, min2 := math.MaxInt, math.MaxInt
    for i := 0; i < len(nums2); i++ {
        if nums1[i] < min1 { min1 = nums1[i] }
        if nums2[i] < min2 { min2 = nums2[i] }
    }
    
    return min2-min1
}