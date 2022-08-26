package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reorderedPowerOf2(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 0",
			args{n: 1},
			true,
		},
		{
			"test 1",
			args{n: 10},
			false,
		},
		{
			"test 2",
			args{n: 46},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reorderedPowerOf2(tt.args.n), "reorderedPowerOf2(%v)", tt.args.n)
		})
	}
}
