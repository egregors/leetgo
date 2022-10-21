package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_containsNearbyDuplicate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test-1",
			args{
				[]int{1, 2, 3, 1},
				3,
			},
			true,
		},
		{
			"test-2",
			args{
				[]int{1, 0, 1, 1},
				1,
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, containsNearbyDuplicate(tt.args.nums, tt.args.k), "containsNearbyDuplicate(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
