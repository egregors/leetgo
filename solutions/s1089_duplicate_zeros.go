/*
	https://leetcode.com/problems/duplicate-zeros/description/

	Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the
	remaining elements to the right.

	Note that elements beyond the length of the original array are not written. Do the above
	modifications to the input array in place and do not return anything.
*/

package solutions

func duplicateZeros(arr []int) {
	n := len(arr)
	for i := n - 2; i >= 0; i-- {
		if arr[i] == 0 {
			arr[n-1] = 0
			for j := n - 1; j > i; j-- {
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
	}
}
