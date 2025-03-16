package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_repairCars(t *testing.T) {
	type args struct {
		ranks []int
		cars  int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"case 0",
			args{
				[]int{4, 2, 3, 1},
				10,
			},
			16,
		},
		{
			"case 1",
			args{
				[]int{5, 1, 6},
				6,
			},
			16,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, repairCars(tt.args.ranks, tt.args.cars), "repairCars(%v, %v)", tt.args.ranks, tt.args.cars)
		})
	}
}
