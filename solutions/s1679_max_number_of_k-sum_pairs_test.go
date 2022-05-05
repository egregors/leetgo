package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxOperations(t *testing.T) {
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
			name: "Test 1",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    5,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				nums: []int{3, 1, 3, 4, 3},
				k:    6,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxOperations(tt.args.nums, tt.args.k), "maxOperations(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
