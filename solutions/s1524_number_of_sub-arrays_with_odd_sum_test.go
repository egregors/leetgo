package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numOfSubarrays(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{[]int{1, 3, 5}},
			4,
		},
		{
			"test 2",
			args{[]int{2, 4, 6}},
			0,
		},
		{
			"test 3",
			args{[]int{1, 2, 3, 4, 5, 6, 7}},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numOfSubarrays(tt.args.arr), "numOfSubarrays(%v)", tt.args.arr)
		})
	}
}
