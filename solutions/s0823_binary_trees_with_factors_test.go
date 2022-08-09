package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numFactoredBinaryTrees(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"0",
			args{arr: []int{2, 4}},
			3,
		},
		{
			"1",
			args{arr: []int{2, 4, 5, 10}},
			7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numFactoredBinaryTrees(tt.args.arr), "numFactoredBinaryTrees(%v)", tt.args.arr)
		})
	}
}
