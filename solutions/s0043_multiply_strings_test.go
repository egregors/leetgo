package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_multiply(t *testing.T) {
	type args struct {
		num1 string
		num2 string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				num1: "2",
				num2: "3",
			},
			want: "6",
		},
		{
			name: "Example 2",
			args: args{
				num1: "123",
				num2: "456",
			},
			want: "56088",
		},
		{
			name: "Example 3",
			args: args{
				num1: "0",
				num2: "0",
			},
			want: "0",
		},
		{
			name: "Example 4",
			args: args{
				num1: "0",
				num2: "1",
			},
			want: "0",
		},
		{
			name: "Example 6",
			args: args{
				num1: "1",
				num2: "1",
			},
			want: "1",
		},
		{
			name: "Example 7",
			args: args{
				num1: "1",
				num2: "2",
			},
			want: "2",
		},
		{
			name: "Example 8",
			args: args{
				num1: "1",
				num2: "3",
			},
			want: "3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, multiply(tt.args.num1, tt.args.num2), "multiply(%v, %v)", tt.args.num1, tt.args.num2)
		})
	}
}
