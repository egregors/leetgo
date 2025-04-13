package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countGoodNumbers(t *testing.T) {
	type args struct {
		n int64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{n: 1},
			want: 5,
		},
		{
			name: "Test 2",
			args: args{n: 4},
			want: 400,
		},
		{
			name: "Test 3",
			args: args{n: 50},
			want: 564908303,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countGoodNumbers(tt.args.n), "countGoodNumbers(%v)", tt.args.n)
		})
	}
}
