package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOne(t *testing.T) {
	type args struct {
		digits []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 0",
			args{
				[]int{1, 2, 3},
			},
			[]int{1, 2, 4},
		},
		{
			"Test 1",
			args{
				[]int{1, 2, 9},
			},
			[]int{1, 3, 0},
		},
		{
			"Test 2",
			args{
				[]int{9, 9, 9},
			},
			[]int{1, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, plusOne(tt.args.digits), "plusOne(%v)", tt.args.digits)
		})
	}
}
