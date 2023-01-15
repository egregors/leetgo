package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numberOfGoodPaths(t *testing.T) {
	type args struct {
		vals  []int
		edges [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{1, 3, 2, 1, 3},
				[][]int{{0, 1}, {0, 2}, {2, 3}, {2, 4}},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfGoodPaths(tt.args.vals, tt.args.edges), "numberOfGoodPaths(%v, %v)", tt.args.vals, tt.args.edges)
		})
	}
}
