package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_triangleType(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test1",
			args{[]int{3, 3, 3}},
			"equilateral",
		},
		{
			"test2",
			args{[]int{3, 4, 5}},
			"scalene",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, triangleType(tt.args.nums), "triangleType(%v)", tt.args.nums)
		})
	}
}
