package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minCapability(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{[]int{2, 3, 5, 9}, 2},
			5,
		},
		{
			"case 2",
			args{[]int{2, 7, 9, 3, 1}, 2},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minCapability(tt.args.nums, tt.args.k), "minCapability(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
