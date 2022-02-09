package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subsetsWithDup(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test 1",
			args{nums: []int{1, 2, 2}},
			[][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, IsEqualAnyOrderIntsSlices(subsetsWithDup(tt.args.nums), tt.want), "subsetsWithDup(%v)")
		})
	}
}
