package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kWeakestRows(t *testing.T) {
	type args struct {
		mat [][]int
		k   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Test 1",
			args{mat: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 0},
				{1, 0, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{1, 1, 1, 1, 1},
			}, k: 3},
			[]int{2, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kWeakestRows(tt.args.mat, tt.args.k), "kWeakestRows(%v, %v)", tt.args.mat, tt.args.k)
		})
	}
}
