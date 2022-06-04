package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 1",
			args: args{
				n: 4,
			},
			want: [][]string{
				{".Q..", "...Q", "Q...", "..Q."},
				{"..Q.", "Q...", "...Q", ".Q.."},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, solveNQueens(tt.args.n), "solveNQueens(%v)", tt.args.n)
		})
	}
}
