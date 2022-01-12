package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchMatrix2(t *testing.T) {
	type args struct {
		matrix [][]int
		target int
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
					{1, 4, 7, 11, 15},
					{2, 5, 8, 12, 19},
					{3, 6, 9, 16, 22},
					{10, 13, 14, 17, 24},
					{18, 21, 23, 26, 30},
				},
				target: 5,
			},
			want: true,
		},
		{
			name: "Test 1",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 50},
				},
				target: 13,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, searchMatrix2(tt.args.matrix, tt.args.target), "searchMatrix2(%v, %v)", tt.args.matrix, tt.args.target)
		})
	}
}
