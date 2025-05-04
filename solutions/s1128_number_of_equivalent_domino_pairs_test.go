package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numEquivDominoPairs(t *testing.T) {
	type args struct {
		dominoes [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				dominoes: [][]int{
					{1, 2},
					{2, 1},
					{3, 4},
					{5, 6},
				},
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				dominoes: [][]int{
					{1, 2},
					{1, 2},
					{1, 1},
					{1, 2},
					{2, 2},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numEquivDominoPairs(tt.args.dominoes), "numEquivDominoPairs(%v)", tt.args.dominoes)
		})
	}
}
