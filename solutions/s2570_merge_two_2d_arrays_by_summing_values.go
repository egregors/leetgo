/*
	https://leetcode.com/problems/merge-two-2d-arrays-by-summing-values/description

	You are given two 2D integer arrays nums1 and nums2.

		nums1[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.
		nums2[i] = [idi, vali] indicate that the number with the id idi has a value equal to vali.

	Each array contains unique ids and is sorted in ascending order by id.

	Merge the two arrays into one array that is sorted in ascending order by id, respecting the following conditions:

		Only ids that appear in at least one of the two arrays should be included in the resulting array.
		Each id should be included only once and its value should be the sum of the values of this id in the two arrays.
	If the id does not exist in one of the two arrays, then assume its value in that array to be 0.

	Return the resulting array. The returned array must be sorted in ascending order by id.
*/

package solutions

func mergeArrays(nums1 [][]int, nums2 [][]int) [][]int {
	return merge2570(nums1, nums2)
}

func merge2570(xs, ys [][]int) [][]int {
	if len(xs) == 0 {
		return ys
	}

	if len(ys) == 0 {
		return xs
	}

	h1, h2 := xs[0], ys[0]
	if h1[0] == h2[0] {
		return append([][]int{{h1[0], h1[1] + h2[1]}}, merge2570(xs[1:], ys[1:])...)
	}

	if h1[0] < h2[0] {
		return append([][]int{h1}, merge2570(xs[1:], ys)...)
	}

	return append([][]int{h2}, merge2570(xs, ys[1:])...)
}
