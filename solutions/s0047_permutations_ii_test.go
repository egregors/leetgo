package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_permuteUnique(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test 1",
			args: args{
				nums: []int{1, 1, 2},
			},
			want: [][]int{
				{1, 1, 2},
				{1, 2, 1},
				{2, 1, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, IsEqualAnyOrderIntsSlices(permuteUnique(tt.args.nums), tt.want), "permuteUnique(%v)")
		})
	}
}
