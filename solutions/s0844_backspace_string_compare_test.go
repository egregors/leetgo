package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_backspaceCompare(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				s: "ab#c",
				t: "ad#c",
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				s: "ab##",
				t: "c#d#",
			},
			want: true,
		},
		{
			name: "Test 3",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
		{
			name: "Test 4",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "Test 5",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
		{
			name: "Test 6",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "Test 7",
			args: args{
				s: "a#c",
				t: "b",
			},
			want: false,
		},
		{
			name: "Test 8",
			args: args{
				s: "a##c",
				t: "#a#c",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, backspaceCompare(tt.args.s, tt.args.t), "backspaceCompare(%v, %v)", tt.args.s, tt.args.t)
		})
	}
}
