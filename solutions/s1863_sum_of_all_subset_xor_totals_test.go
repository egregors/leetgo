package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetXORSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test_case_1",
			args: args{
				nums: []int{1, 3},
			},
			want: 6,
		},
		{
			name: "test_case_2",
			args: args{
				nums: []int{5, 1, 6},
			},
			want: 28,
		},
		{
			name: "test_case_3",
			args: args{
				nums: []int{3, 4, 5, 6, 7, 8},
			},
			want: 480,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, subsetXORSum(tt.args.nums), "subsetXORSum(%v)", tt.args.nums)
		})
	}
}
