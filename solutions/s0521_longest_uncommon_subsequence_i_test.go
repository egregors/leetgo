package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findLUSlength(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				a: "aba",
				b: "cdc",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				a: "aaa",
				b: "aaa",
			},
			want: -1,
		},
		{
			name: "Test 5",
			args: args{
				a: "aaa",
				b: "",
			},
			want: 3,
		},
		{
			name: "Test 6",
			args: args{
				a: "",
				b: "aaa",
			},
			want: 3,
		},
		{
			name: "Test 7",
			args: args{
				a: "",
				b: "",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findLUSlength(tt.args.a, tt.args.b), "findLUSlength(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
