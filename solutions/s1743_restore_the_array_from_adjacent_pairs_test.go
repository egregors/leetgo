package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_restoreArray(t *testing.T) {
	type args struct {
		adjacentPairs [][]int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"case 1",
			args{[][]int{{2, 1}, {3, 4}, {3, 2}}},
			[]int{4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, IsEqualAnyOrderInts(tt.want, restoreArray(tt.args.adjacentPairs)), "restoreArray(%v)", tt.args.adjacentPairs)
		})
	}
}
