package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numberOfAlternatingGroups(t *testing.T) {
	type args struct {
		colors []int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]int{0, 1, 0, 1, 0},
				3,
			},
			3,
		},
		{
			"case 2",
			args{
				[]int{0, 1, 0, 0, 1, 0, 1},
				6,
			},
			2,
		},
		{
			"case 3",
			args{
				[]int{1, 1, 0, 1},
				4,
			},
			0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numberOfAlternatingGroups(tt.args.colors, tt.args.k), "numberOfAlternatingGroups(%v, %v)", tt.args.colors, tt.args.k)
		})
	}
}
