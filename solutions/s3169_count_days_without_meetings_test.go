package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countDays(t *testing.T) {
	type args struct {
		days     int
		meetings [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"case 1",
			args{
				10,
				[][]int{
					{5, 7},
					{1, 3},
					{9, 10},
				},
			},
			2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, countDays(tt.args.days, tt.args.meetings), "countDays(%v, %v)", tt.args.days, tt.args.meetings)
		})
	}
}
