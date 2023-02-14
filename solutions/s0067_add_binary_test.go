package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_addBinary(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test 1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "test 2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addBinary(tt.args.a, tt.args.b), "addBinary(%v, %v)", tt.args.a, tt.args.b)
		})
	}
}
