package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPairs2176(t *testing.T) {
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
			name: "Test1",
			args: args{
				nums: []int{3, 1, 2, 2, 2, 1, 3},
				k:    2,
			},
			want: 4,
		},
		{
			name: "Test2",
			args: args{
				nums: []int{1, 2, 3, 4},
				k:    1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countPairs2176(tt.args.nums, tt.args.k), "countPairs2176(%v, %v)", tt.args.nums, tt.args.k)
		})
	}
}
