package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestNumber3370(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{n: 5},
			want: 7,
		},
		{
			name: "test case 2",
			args: args{n: 10},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, smallestNumber3370(tt.args.n), "smallestNumber3370(%v)", tt.args.n)
		})
	}
}
