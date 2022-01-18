package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test 1",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1},
				n:         1,
			},
			want: true,
		},
		{
			name: "Test 2",
			args: args{
				flowerbed: []int{1, 0, 0, 0, 1, 0, 0},
				n:         2,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, canPlaceFlowers(tt.args.flowerbed, tt.args.n), "canPlaceFlowers(%v, %v)", tt.args.flowerbed, tt.args.n)
		})
	}
}
