package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pushDominoes(t *testing.T) {
	type args struct {
		dominoes string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				dominoes: "RR.L",
			},
			want: "RR.L",
		},
		{
			name: "Test 2",
			args: args{
				dominoes: ".L.R...LR..L..",
			},
			want: "LL.RR.LLRRLL..",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, pushDominoes(tt.args.dominoes), "pushDominoes(%v)", tt.args.dominoes)
		})
	}
}
