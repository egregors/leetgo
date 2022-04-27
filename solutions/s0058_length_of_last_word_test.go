package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lengthOfLastWord(t *testing.T) {
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
				s: "Hello World",
			},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{
				s: "a ",
			},
			want: 1,
		},
		{
			name: "Test 3",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				s: " ",
			},
			want: 0,
		},
		{
			name: "Test 5",
			args: args{
				s: "",
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lengthOfLastWord(tt.args.s), "lengthOfLastWord(%v)", tt.args.s)
		})
	}
}
