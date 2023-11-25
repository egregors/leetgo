package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getSumAbsoluteDifferences(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 0",
			args{[]int{2, 3, 5}},
			[]int{4, 3, 5},
		},
		{
			"case 1",
			args{[]int{1, 4, 6, 8, 10}},
			[]int{24, 15, 13, 15, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, getSumAbsoluteDifferences(tt.args.nums), "getSumAbsoluteDifferences(%v)", tt.args.nums)
		})
	}
}
