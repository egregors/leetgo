package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mirrorReflection(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test0",
			args{p: 2, q: 1},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, mirrorReflection(tt.args.p, tt.args.q), "mirrorReflection(%v, %v)", tt.args.p, tt.args.q)
		})
	}
}
