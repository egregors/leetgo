package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_eraseOverlapIntervals(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{intervals: [][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}}},
			1,
		},
		{
			"Test 2",
			args{intervals: [][]int{{1, 2}, {1, 2}, {1, 2}}},
			2,
		},
		{
			"Test 3",
			args{intervals: [][]int{{1, 2}, {2, 3}}},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, eraseOverlapIntervals(tt.args.intervals), "eraseOverlapIntervals(%v)", tt.args.intervals)
		})
	}
}
