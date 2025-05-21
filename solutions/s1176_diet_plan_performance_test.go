package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_dietPlanPerformance(t *testing.T) {
	type args struct {
		calories []int
		k        int
		lower    int
		upper    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				calories: []int{1, 2, 3, 4, 5},
				k:        1,
				lower:    3,
				upper:    3,
			},
			0,
		},
		{
			"test 2",
			args{
				calories: []int{3, 2},
				k:        2,
				lower:    0,
				upper:    1,
			},
			1,
		},
		{
			"test 3",
			args{
				calories: []int{6, 5, 0, 0},
				k:        2,
				lower:    1,
				upper:    5,
			},
			0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, dietPlanPerformance(tt.args.calories, tt.args.k, tt.args.lower, tt.args.upper), "dietPlanPerformance(%v, %v, %v, %v)", tt.args.calories, tt.args.k, tt.args.lower, tt.args.upper)
		})
	}
}
