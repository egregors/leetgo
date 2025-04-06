package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestDivisibleSubset(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				nums: []int{1, 2, 3},
			},
			want: []int{1, 2},
		},
		{
			name: "test2",
			args: args{
				nums: []int{1, 2, 4, 8},
			},
			want: []int{1, 2, 4, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, largestDivisibleSubset(tt.args.nums), "largestDivisibleSubset(%v)", tt.args.nums)
		})
	}
}
