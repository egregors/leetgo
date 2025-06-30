package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLHS(t *testing.T) {
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
			args: args{nums: []int{1, 3, 2, 2, 5, 2, 3, 7}},
			want: 5,
		},
		{
			name: "Example 2",
			args: args{nums: []int{1, 2, 3, 4}},
			want: 2,
		},
		{
			name: "Example 3",
			args: args{nums: []int{1, 1, 1, 1}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLHS(tt.args.nums), "findLHS(%v)", tt.args.nums)
		})
	}
}
