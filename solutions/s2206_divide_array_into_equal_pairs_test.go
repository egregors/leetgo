package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divideArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"case 1",
			args{[]int{3, 2, 3, 2, 2, 2}},
			true,
		},
		{
			"case 2",
			args{[]int{1, 2, 3, 4}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, divideArray(tt.args.nums), "divideArray(%v)", tt.args.nums)
		})
	}
}
