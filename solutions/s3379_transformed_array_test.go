package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_constructTransformedArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{3, -2, 1, 1}},
			want: []int{1, 1, 1, 3},
		},
		{
			name: "Example 2",
			args: args{nums: []int{-1, 4, -1}},
			want: []int{-1, -1, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, constructTransformedArray(tt.args.nums), "constructTransformedArray(%v)", tt.args.nums)
		})
	}
}
