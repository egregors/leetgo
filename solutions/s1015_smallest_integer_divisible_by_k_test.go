package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestRepunitDivByK(t *testing.T) {
	tests := []struct {
		name string
		k    int
		want int
	}{
		{
			name: "test case 1",
			k:    1,
			want: 1,
		},
		{
			name: "test case 2",
			k:    2,
			want: -1,
		},
		{
			name: "test case 3",
			k:    3,
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, smallestRepunitDivByK(tt.k), "smallestRepunitDivByK(%v)", tt.k)
		})
	}
}
