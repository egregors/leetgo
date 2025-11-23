package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSumDivThree(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "example 1",
			args: args{nums: []int{3, 6, 5, 1, 8}},
			want: 18,
		},
		{
			name: "example 2",
			args: args{nums: []int{4}},
			want: 0,
		},
		{
			name: "example 3",
			args: args{nums: []int{1, 2, 3, 4, 4}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxSumDivThree(tt.args.nums), "maxSumDivThree(%v)", tt.args.nums)
		})
	}
}
