package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maximumLength(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{[]int{1, 2, 3, 4}},
			4,
		},
		{
			"test2",
			args{[]int{1, 2, 1, 1, 2, 1, 2}},
			6,
		},
		{
			"test3",
			args{[]int{1, 3}},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumLength(tt.args.nums), "maximumLength(%v)", tt.args.nums)
		})
	}
}
