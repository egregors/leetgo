package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isHappy(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"test 1",
			args{n: 19},
			true,
		},
		{
			"test 2",
			args{n: 2},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, isHappy(tt.args.n), "isHappy(%v)", tt.args.n)
		})
	}
}
