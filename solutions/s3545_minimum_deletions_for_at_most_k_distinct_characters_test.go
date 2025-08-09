package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minDeletion(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				s: "abc",
				k: 2,
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				s: "aabb",
				k: 2,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minDeletion(tt.args.s, tt.args.k), "minDeletion(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
