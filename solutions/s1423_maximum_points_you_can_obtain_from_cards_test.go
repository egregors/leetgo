package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxScore(t *testing.T) {
	type args struct {
		cardPoints []int
		k          int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{[]int{1, 2, 3, 4, 5, 6, 1}, 3},
			12,
		},
		{
			"Test 1",
			args{[]int{2, 2, 2}, 2},
			4,
		},
		{
			"Test 2",
			args{[]int{9, 7, 7, 9, 7, 7, 9}, 7},
			55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxScore(tt.args.cardPoints, tt.args.k), "maxScore(%v, %v)", tt.args.cardPoints, tt.args.k)
		})
	}
}
