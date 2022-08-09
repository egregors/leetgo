package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kthSmallest378(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, k: 8},
			13,
		},
		{
			"1",
			args{matrix: [][]int{{-5}}, k: 1},
			-5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kthSmallest378(tt.args.matrix, tt.args.k), "kthSmallest(%v, %v)", tt.args.matrix, tt.args.k)
		})
	}
}
