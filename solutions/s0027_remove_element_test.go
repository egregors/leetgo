package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_removeElement(t *testing.T) {
	type args struct {
		nums []int
		val  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				nums: []int{3, 2, 2, 3},
				val:  3,
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{0, 1, 2, 2, 3, 0, 4, 2},
				val:  2,
			},
			want: 5,
		},
		{
			name: "Example 3",
			args: args{
				nums: []int{1},
				val:  1,
			},
			want: 0,
		},
		{
			name: "Example 4",
			args: args{
				nums: []int{4, 5},
				val:  4,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, removeElement(tt.args.nums, tt.args.val), "removeElement(%v, %v)", tt.args.nums, tt.args.val)
		})
	}
}
