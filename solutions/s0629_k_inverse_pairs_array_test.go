package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_kInversePairs(t *testing.T) {
	type args struct {
		n int
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 0",
			args{3, 0},
			1,
		},
		{
			"test 1",
			args{3, 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, kInversePairs(tt.args.n, tt.args.k), "kInversePairs(%v, %v)", tt.args.n, tt.args.k)
		})
	}
}
