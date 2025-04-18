package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_judgeCircle(t *testing.T) {
	type args struct {
		moves string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				moves: "UD",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				moves: "LL",
			},
			want: false,
		},
		{
			name: "Test 3",
			args: args{
				moves: "RRDD",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, judgeCircle(tt.args.moves), "judgeCircle(%v)", tt.args.moves)
		})
	}
}
