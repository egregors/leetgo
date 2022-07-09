package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxResult(t *testing.T) {
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
			name: "test1",
			args: args{
				nums: []int{1, -1, -2, 4, -7, 3},
				k:    2,
			},
			want: 7,
		},
		{
			name: "test2",
			args: args{
				nums: []int{10, -5, -2, 4, 0, 3},
				k:    3,
			},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxResult(tt.args.nums, tt.args.k), "maxResult(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
