package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimizeArrayValue(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{3, 7, 1, 6},
			},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimizeArrayValue(tt.args.nums), "minimizeArrayValue(%v)", tt.args.nums)
		})
	}
}
