package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countVowelPermutation(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example1",
			args: args{
				n: 1,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countVowelPermutation(tt.args.n), "countVowelPermutation(%v)", tt.args.n)
		})
	}
}
