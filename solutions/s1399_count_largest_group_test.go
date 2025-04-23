package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countLargestGroup(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{n: 13},
			want: 4,
		},
		{
			name: "Test 2",
			args: args{n: 2},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countLargestGroup(tt.args.n), "countLargestGroup(%v)", tt.args.n)
		})
	}
}
