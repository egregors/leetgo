package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumRounds(t *testing.T) {
	type args struct {
		tasks []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{2, 2, 3, 3, 2, 4, 4, 4, 4, 4},
			},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumRounds(tt.args.tasks), "minimumRounds(%v)", tt.args.tasks)
		})
	}
}
