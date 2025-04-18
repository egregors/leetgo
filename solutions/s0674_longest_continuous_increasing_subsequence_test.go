package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLengthOfLCIS(t *testing.T) {
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
				nums: []int{1, 3, 5, 4, 7},
			},
			want: 3,
		},
		{
			name: "test_case_2",
			args: args{
				nums: []int{2, 2, 2, 2, 2},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLengthOfLCIS(tt.args.nums), "findLengthOfLCIS(%v)", tt.args.nums)
		})
	}
}
