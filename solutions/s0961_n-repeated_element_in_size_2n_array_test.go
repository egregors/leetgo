package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_repeatedNTimes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"1",
			args{[]int{1, 2, 3, 3}},
			3,
		},
		{
			"2",
			args{[]int{2, 1, 2, 5, 3, 2}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, repeatedNTimes(tt.args.nums), "repeatedNTimes(%v)", tt.args.nums)
		})
	}
}
