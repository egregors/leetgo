package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_calculate(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "3+2*2",
			},
			want: 7,
		},
		{
			name: "Test 2",
			args: args{
				s: " 3/2 ",
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				s: " 3+5 / 2",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, calculate(tt.args.s), "calculate(%v)", tt.args.s)
		})
	}
}
