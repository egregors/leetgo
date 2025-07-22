package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumUniqueSubarray1695(t *testing.T) {
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
			args: args{nums: []int{4, 2, 4, 5, 6}},
			want: 17,
		}, {
			name: "Example 2",
			args: args{nums: []int{5, 2, 1, 2, 5, 2, 1, 2, 5}},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumUniqueSubarray1695(tt.args.nums), "maximumUniqueSubarray1695(%v)", tt.args.nums)
		})
	}
}
