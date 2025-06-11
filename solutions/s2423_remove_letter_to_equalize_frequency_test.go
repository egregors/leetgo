package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_equalFrequency(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				word: "abcc",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				word: "aazz",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, equalFrequency(tt.args.word), "equalFrequency(%v)", tt.args.word)
		})
	}
}
