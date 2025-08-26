package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_areaOfMaxDiagonal(t *testing.T) {
	type args struct {
		dimensions [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{dimensions: [][]int{{9, 3}, {8, 6}}},
			want: 48,
		},
		{
			name: "test2",
			args: args{dimensions: [][]int{{3, 4}, {4, 3}}},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, areaOfMaxDiagonal(tt.args.dimensions), "areaOfMaxDiagonal(%v)", tt.args.dimensions)
		})
	}
}
