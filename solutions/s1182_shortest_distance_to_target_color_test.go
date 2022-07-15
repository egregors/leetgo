package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_shortestDistanceColor(t *testing.T) {
	type args struct {
		colors  []int
		queries [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example1",
			args: args{
				colors:  []int{1, 1, 2, 1, 3, 2, 2, 3, 3},
				queries: [][]int{{1, 3}, {2, 2}, {6, 1}},
			},
			want: []int{3, 0, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, shortestDistanceColor(tt.args.colors, tt.args.queries), "shortestDistanceColor(%v, %v)", tt.args.colors, tt.args.queries)
		})
	}
}
