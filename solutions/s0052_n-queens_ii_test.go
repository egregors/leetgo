package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_totalNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 0",
			args{n: 4},
			2,
		},
		{
			"Test 1",
			args{n: 1},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, totalNQueens(tt.args.n), "totalNQueens(%v)", tt.args.n)
		})
	}
}
