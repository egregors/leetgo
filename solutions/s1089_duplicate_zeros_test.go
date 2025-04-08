/*
	https://leetcode.com/problems/duplicate-zeros/description/

	Given a fixed-length integer array arr, duplicate each occurrence of zero, shifting the
	remaining elements to the right.

	Note that elements beyond the length of the original array are not written. Do the above
	modifications to the input array in place and do not return anything.
*/

package solutions

import "testing"

func Test_duplicateZeros(t *testing.T) {
	tests := []struct {
		name string
		arr  []int
		want []int
	}{
		{
			name: "Test 1",
			arr:  []int{1, 0, 2, 3, 0, 4, 5, 0},
			want: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			name: "Test 2",
			arr:  []int{1, 2, 3},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duplicateZeros(tt.arr)
			if len(tt.arr) != len(tt.want) {
				t.Errorf("duplicateZeros() = %v, want %v", tt.arr, tt.want)
			}
			for i := range tt.arr {
				if tt.arr[i] != tt.want[i] {
					t.Errorf("duplicateZeros() = %v, want %v", tt.arr, tt.want)
					break
				}
			}
		})
	}
}
