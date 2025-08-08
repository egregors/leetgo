package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findMiddleIndex(t *testing.T) {
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
			args: args{nums: []int{2, 3, -1, 8, 4}},
			want: 3,
		}, {
			name: "Example 2",
			args: args{nums: []int{1, -1, 4}},
			want: 2,
		}, {
			name: "Example 3",
			args: args{nums: []int{2, 5}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMiddleIndex(tt.args.nums), "findMiddleIndex(%v)", tt.args.nums)
		})
	}
}
