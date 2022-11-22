package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_diStringMatch(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Example 1",
			args: args{
				s: "IDID",
			},
			want: []int{0, 4, 1, 3, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, diStringMatch(tt.args.s), "diStringMatch(%v)", tt.args.s)
		})
	}
}
