/*
	https://leetcode.com/problems/type-of-triangle/description/

		You are given a 0-indexed integer array nums of size 3 which can form the
			sides of a triangle.

		A triangle is called equilateral if it has all sides of equal length.
		A triangle is called isosceles if it has exactly two sides of equal length.
		A triangle is called scalene if all its sides are of different lengths.

	Return a string representing the type of triangle that can be formed or "none"
		if it cannot form a triangle.
*/

package solutions

import "sort"

func triangleType(nums []int) string {
	canFormed := func(xs []int) bool {
		sort.Ints(xs)
		return (nums[0] + nums[1]) > nums[2]
	}

	a, b, c := nums[0], nums[1], nums[2]
	if canFormed(nums) {
		if a == b && b == c {
			return "equilateral"
		}
		if a == b || b == c || c == a {
			return "isosceles"
		}
		return "scalene"
	}

	return "none"
}
