package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumAverageDifference(t *testing.T) {
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
				[]int{2, 5, 3, 9, 5, 3},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumAverageDifference(tt.args.nums), "minimumAverageDifference(%v)", tt.args.nums)
		})
	}
}
