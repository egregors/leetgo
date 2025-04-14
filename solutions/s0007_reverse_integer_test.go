package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse0007(t *testing.T) {
	tests := []struct {
		name string
		x    int
		want int
	}{
		{
			name: "test1",
			x:    123,
			want: 321,
		},
		{
			name: "test2",
			x:    -123,
			want: -321,
		},
		{
			name: "test3",
			x:    120,
			want: 21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverse0007(tt.x), "reverse0007(%v)", tt.x)
		})
	}
}
