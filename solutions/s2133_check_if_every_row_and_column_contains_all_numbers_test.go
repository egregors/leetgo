package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_checkValid(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				matrix: [][]int{
					{1, 2, 3},
					{3, 1, 2},
					{2, 3, 1},
				},
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				matrix: [][]int{
					{1, 1, 1},
					{1, 2, 3},
					{1, 2, 3},
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkValid(tt.args.matrix), "checkValid(%v)", tt.args.matrix)
		})
	}
}
