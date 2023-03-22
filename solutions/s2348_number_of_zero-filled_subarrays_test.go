package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_zeroFilledSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"1",
			args{[]int{1, 3, 0, 0, 2, 0, 0, 4}},
			6,
		},
		{
			"2",
			args{[]int{0, 0, 0, 2, 0, 0}},
			9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, zeroFilledSubarray(tt.args.nums), "zeroFilledSubarray(%v)", tt.args.nums)
		})
	}
}
