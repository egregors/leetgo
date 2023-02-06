package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_shuffle(t *testing.T) {
	type args struct {
		nums []int
		n    int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Example 1",
			args{
				[]int{2, 5, 1, 3, 4, 7},
				3,
			},
			[]int{2, 3, 5, 4, 1, 7},
		},
		{
			"Example 2",
			args{
				[]int{1, 2, 3, 4, 4, 3, 2, 1},
				4,
			},
			[]int{1, 4, 2, 3, 3, 2, 4, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shuffle(tt.args.nums, tt.args.n), "shuffle(%v, %v)", tt.args.nums, tt.args.n)
		})
	}
}
