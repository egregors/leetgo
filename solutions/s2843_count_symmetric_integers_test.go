package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countSymmetricIntegers(t *testing.T) {
	type args struct {
		low  int
		high int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{low: 1, high: 100},
			want: 9,
		},
		{
			name: "Test Case 2",
			args: args{low: 1200, high: 1230},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countSymmetricIntegers(tt.args.low, tt.args.high), "countSymmetricIntegers(%v, %v)", tt.args.low, tt.args.high)
		})
	}
}
