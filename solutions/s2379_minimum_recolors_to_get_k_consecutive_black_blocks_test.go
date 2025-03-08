package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minimumRecolors(t *testing.T) {
	type args struct {
		blocks string
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{"WBBWWBBWBW", 7},
			3,
		},
		{
			"test 2",
			args{"WBWBBBW", 2},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumRecolors(tt.args.blocks, tt.args.k), "minimumRecolors(%v, %v)", tt.args.blocks, tt.args.k)
		})
	}
}
