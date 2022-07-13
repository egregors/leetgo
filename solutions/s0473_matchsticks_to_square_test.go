package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_makesquare(t *testing.T) {
	type args struct {
		matchsticks []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Test 0",
			args{
				[]int{1, 1, 2, 2, 2},
			},
			true,
		},
		{
			"Test 1",
			args{
				[]int{3, 3, 3, 3, 4},
			},
			false,
		},
		{
			"test 2",
			args{
				[]int{5, 5, 5, 5, 4, 4, 4, 4, 3, 3, 3, 3},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, makesquare(tt.args.matchsticks), "makesquare(%v)", tt.args.matchsticks)
		})
	}
}
