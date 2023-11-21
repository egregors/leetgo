package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countNicePairs(t *testing.T) {
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
			args{[]int{42, 11, 1, 97}},
			2,
		},
		{
			"case 2",
			args{[]int{13, 10, 35, 24, 76}},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countNicePairs(tt.args.nums), "countNicePairs(%v)", tt.args.nums)
		})
	}
}
