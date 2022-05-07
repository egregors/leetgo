package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_find132pattern(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 1",
			args{nums: []int{1, 2, 3, 4}},
			false,
		},
		{
			"Test 2",
			args{nums: []int{3, 1, 4, 2}},
			true,
		}, {
			"Test 3",
			args{nums: []int{-1, 3, 2, 0}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, find132pattern(tt.args.nums), "find132pattern(%v)", tt.args.nums)
		})
	}
}
