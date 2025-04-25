/*
	2657. Find the Prefix Common Array of Two Arrays

	You are given two 0-indexed integer permutations A and B of length n.

	A prefix common array of A and B is an array C such that C[i] is equal to the count of numbers that are
	present at or before the index i in both A and B.

	Return the prefix common array of A and B.

	A sequence of n integers is called a permutation if it contains all integers from 1 to n exactly once.
*/

package solutions

import (
	"reflect"
	"testing"
)

func Test_findThePrefixCommonArray(t *testing.T) {
	type args struct {
		A []int
		B []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				A: []int{1, 3, 2, 4},
				B: []int{3, 1, 2, 4},
			},
			want: []int{0, 2, 3, 4},
		},
		{
			name: "test2",
			args: args{
				A: []int{2, 3, 1},
				B: []int{3, 1, 2},
			},
			want: []int{0, 1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findThePrefixCommonArray(tt.args.A, tt.args.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findThePrefixCommonArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
