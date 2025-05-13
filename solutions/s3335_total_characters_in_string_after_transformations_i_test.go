package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthAfterTransformations(t *testing.T) {
	type args struct {
		s string
		t int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{"abcyy", 2},
			7,
		},
		{
			"Test 2",
			args{"azbk", 5},
			5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthAfterTransformations(tt.args.s, tt.args.t), "lengthAfterTransformations(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
