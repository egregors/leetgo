package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_longestCycle(t *testing.T) {
	type args struct {
		edges []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{edges: []int{3, 3, 4, 2, 3}},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, longestCycle(tt.args.edges), "longestCycle(%v)", tt.args.edges)
		})
	}
}
