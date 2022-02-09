package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"Test 1",
			args{candidates: []int{2, 3, 6, 7}, target: 7},
			[][]int{{2, 2, 3}, {7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, IsEqualAnyOrderIntsSlices(combinationSum(tt.args.candidates, tt.args.target), tt.want), "combinationSum(%v, %v)")
		})
	}
}
