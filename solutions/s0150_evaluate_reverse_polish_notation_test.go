package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_evalRPN(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test1",
			args: args{
				tokens: []string{"2", "1", "+", "3", "*"},
			},
			want: 9,
		},
		{
			name: "Test2",
			args: args{
				tokens: []string{"4", "13", "5", "/", "+"},
			},
			want: 6,
		},
		{
			name: "Test3",
			args: args{
				tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, evalRPN(tt.args.tokens), "evalRPN(%v)", tt.args.tokens)
		})
	}
}
