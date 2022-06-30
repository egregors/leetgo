package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minMoves2(t *testing.T) {
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
			args: args{
				nums: []int{1, 2, 3},
			},
			want: 2,
		},
		{
			name: "Example 2",
			args: args{
				nums: []int{1, 10, 2, 9},
			},
			want: 16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minMoves2(tt.args.nums), "minMoves2(%v)", tt.args.nums)
		})
	}
}
