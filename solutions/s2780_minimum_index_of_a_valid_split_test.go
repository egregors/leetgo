package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumIndex(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{1, 2, 2, 2}},
			2,
		},
		{
			"case 2",
			args{[]int{2, 1, 3, 1, 1, 1, 7, 1, 2, 1}},
			4,
		},
		{
			"case 3",
			args{[]int{3, 3, 3, 3, 7, 2, 2}},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumIndex(tt.args.nums), "minimumIndex(%v)", tt.args.nums)
		})
	}
}
