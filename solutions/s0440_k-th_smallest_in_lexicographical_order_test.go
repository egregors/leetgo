package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthNumber440(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		//Example 1:
		//
		//Input: n = 13, k = 2
		//Output: 10
		//Explanation: The lexicographical order is [1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9], so the second smallest number is 10.
		//
		//Example 2:
		//
		//Input: n = 1, k = 1
		//Output: 1
		{
			name: "Example 1",
			args: args{
				n: 13,
				k: 2,
			},
			want: 10,
		},
		{
			name: "Example 2",
			args: args{
				n: 1,
				k: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findKthNumber440(tt.args.n, tt.args.k), "findKthNumber440(%v, %v)", tt.args.n, tt.args.k)
		})
	}
}
