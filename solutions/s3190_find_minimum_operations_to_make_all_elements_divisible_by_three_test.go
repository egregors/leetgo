package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumOperations3190(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 3,
		},
		{
			name: "Example 2",
			args: args{nums: []int{3, 6, 9}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumOperations3190(tt.args.nums), "minimumOperations3190(%v)", tt.args.nums)
		})
	}
}
