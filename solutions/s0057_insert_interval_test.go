package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_insert57(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			"test-1",
			args{
				[][]int{
					{1, 3},
					{6, 9},
				},
				[]int{2, 5},
			},
			[][]int{
				{1, 5},
				{6, 9},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, insert57(tt.args.intervals, tt.args.newInterval), "insert57(%v, %v)", tt.args.intervals, tt.args.newInterval)
		})
	}
}
