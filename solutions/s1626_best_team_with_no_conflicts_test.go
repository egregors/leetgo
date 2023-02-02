package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bestTeamScore(t *testing.T) {
	type args struct {
		scores []int
		ages   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test 1",
			args{
				[]int{1, 3, 5, 10, 15},
				[]int{1, 2, 3, 4, 5},
			},
			34,
		},
		{
			"test 2",
			args{
				[]int{4, 5, 6, 5},
				[]int{2, 1, 2, 1},
			},
			16,
		},
		{
			"test 3",
			args{
				[]int{1, 2, 3, 5},
				[]int{8, 9, 10, 1},
			},
			6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, bestTeamScore(tt.args.scores, tt.args.ages), "bestTeamScore(%v, %v)", tt.args.scores, tt.args.ages)
		})
	}
}
