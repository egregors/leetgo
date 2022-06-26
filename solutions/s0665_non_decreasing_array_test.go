package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkPossibility(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 0",
			args{nums: []int{4, 2, 3}},
			true,
		},
		{
			"Test 1",
			args{nums: []int{4, 2, 1}},
			false,
		},
		{
			"Test 2",
			args{nums: []int{3, 4, 2, 3}},
			false,
		},
		{
			"Test 3",
			args{nums: []int{5, 7, 1, 8}},
			true,
		},
		{
			"Test 4",
			args{nums: []int{-1, 4, 2, 3}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkPossibility(tt.args.nums), "checkPossibility(%v)", tt.args.nums)
		})
	}
}
