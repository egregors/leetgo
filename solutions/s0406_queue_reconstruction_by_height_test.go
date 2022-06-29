package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reconstructQueue(t *testing.T) {
	type args struct {
		people [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Example 1",
			args: args{
				people: [][]int{
					{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2},
				},
			},
			want: [][]int{
				{5, 0}, {7, 0}, {5, 2}, {6, 1}, {4, 4}, {7, 1},
			},
		},
		{
			name: "Example 2",
			args: args{
				people: [][]int{
					{6, 0}, {5, 0}, {4, 0}, {3, 2}, {2, 2}, {1, 4},
				},
			},
			want: [][]int{
				{4, 0}, {5, 0}, {2, 2}, {3, 2}, {1, 4}, {6, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reconstructQueue(tt.args.people), "reconstructQueue(%v)", tt.args.people)
		})
	}
}
