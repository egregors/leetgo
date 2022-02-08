package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_addDigits(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				num: 38,
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				num: 9,
			},
			want: 9,
		},
		{
			name: "Test 3",
			args: args{
				num: 10,
			},
			want: 1,
		},
		{
			name: "Test 4",
			args: args{
				num: 0,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, addDigits(tt.args.num), "addDigits(%v)", tt.args.num)
		})
	}
}
