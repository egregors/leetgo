package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximumBags(t *testing.T) {
	type args struct {
		capacity        []int
		rocks           []int
		additionalRocks int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				[]int{2, 3, 4, 5},
				[]int{1, 2, 4, 4},
				2,
			},
			3,
		},
		{
			"case 2",
			args{
				[]int{10, 2, 2},
				[]int{2, 2, 0},
				100,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximumBags(tt.args.capacity, tt.args.rocks, tt.args.additionalRocks), "maximumBags(%v, %v, %v)", tt.args.capacity, tt.args.rocks, tt.args.additionalRocks)
		})
	}
}
