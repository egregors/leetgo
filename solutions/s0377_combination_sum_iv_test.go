package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum4(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 0",
			args{nums: []int{1, 2, 3}, target: 4},
			7,
		},
		{
			"test 1",
			args{nums: []int{9}, target: 3},
			0,
		},
		{
			"test 2",
			args{nums: []int{2, 1, 3}, target: 35},
			1132436852,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, combinationSum4(tt.args.nums, tt.args.target), "combinationSum4(%v, %v)", tt.args.nums, tt.args.target)
		})
	}
}
