package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaximizedCapital(t *testing.T) {
	type args struct {
		k       int
		w       int
		profits []int
		capital []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				2,
				0,
				[]int{1, 2, 3},
				[]int{0, 1, 1},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findMaximizedCapital(tt.args.k, tt.args.w, tt.args.profits, tt.args.capital), "findMaximizedCapital(%v, %v, %v, %v)", tt.args.k, tt.args.w, tt.args.profits, tt.args.capital)
		})
	}
}
