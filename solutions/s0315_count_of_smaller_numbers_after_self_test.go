package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSmaller(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"test 0",
			args{[]int{5, 2, 6, 1}},
			[]int{2, 1, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSmaller(tt.args.nums), "countSmaller(%v)", tt.args.nums)
		})
	}
}
