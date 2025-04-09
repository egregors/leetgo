package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minOperations3375(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				nums: []int{5, 2, 5, 4, 5},
				k:    2,
			},
			want: 2,
		},
		{
			name: "Test Case 2",
			args: args{
				nums: []int{2, 1, 2},
				k:    2,
			},
			want: -1,
		},
		{
			name: "Test Case 3",
			args: args{
				nums: []int{9, 7, 5, 3},
				k:    1,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minOperations3375(tt.args.nums, tt.args.k), "minOperations3375(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
