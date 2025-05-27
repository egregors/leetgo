/*
	https://leetcode.com/problems/minimum-recolors-to-get-k-consecutive-black-blocks/description/

	You are given a 0-indexed string blocks of length n, where blocks[i] is either
		'W' or 'B', representing the color
	of the ith block. The characters 'W' and 'B' denote the colors white and black,
		respectively.

	You are also given an integer k, which is the desired number of consecutive
		black blocks.

	In one operation, you can recolor a white block such that it becomes a black
		block.

	Return the minimum number of operations needed such that there is at least one
		occurrence of k consecutive
	black blocks.
*/

package solutions

import "math"

func minimumRecolors(blocks string, k int) int {
	xs := make([]int, len(blocks)-k+1)
	for i := 0; i < k; i++ {
		if blocks[i] == 'W' {
			xs[0]++
		}
	}

	for i := 1; i <= len(blocks)-k; i++ {
		xs[i] = xs[i-1]
		if blocks[i-1] == 'W' {
			xs[i]--
		}
		if blocks[i+k-1] == 'W' {
			xs[i]++
		}
	}

	res := math.MaxInt
	for _, x := range xs {
		if x < res {
			res = x
		}
	}

	return res
}
