package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findPaths(t *testing.T) {
	type args struct {
		m           int
		n           int
		maxMove     int
		startRow    int
		startColumn int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 0",
			args{2, 2, 2, 0, 0},
			6,
		},
		{
			"test 1",
			args{1, 3, 3, 0, 1},
			12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findPaths(tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn), "findPaths(%v, %v, %v, %v, %v)", tt.args.m, tt.args.n, tt.args.maxMove, tt.args.startRow, tt.args.startColumn)
		})
	}
}
