/*
	https://leetcode.com/problems/container-with-most-water/

	You are given an integer array height of length n. There are n vertical lines
		drawn such
	that the two endpoints of the ith line are (i, 0) and (i, height[i]).

	Find two lines that together with the x-axis form a container, such that the
		container
	contains the most water.

	Return the maximum amount of water a container can store.

	Notice that you may not slant the container.
*/

package solutions

func maxArea(height []int) int {
	res, lo, hi := 0, 0, len(height)-1
	for lo < hi {
		if v := Minimum(height[lo], height[hi]) * (hi - lo); v > res {
			res = v
		}

		if height[lo] <= height[hi] {
			lo++
		} else {
			hi--
		}
	}
	return res
}
