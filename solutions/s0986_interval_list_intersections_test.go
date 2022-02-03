package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_intervalIntersection(t *testing.T) {
	type args struct {
		firstList  [][]int
		secondList [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test1",
			args: args{
				firstList: [][]int{
					{0, 2},
					{5, 10},
					{13, 23},
					{24, 25},
				},
				secondList: [][]int{
					{1, 5},
					{8, 12},
					{15, 24},
					{25, 25},
				},
			},
			want: [][]int{
				{1, 2},
				{5, 5},
				{8, 10},
				{15, 23},
				{24, 24},
				{25, 25},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, intervalIntersection(tt.args.firstList, tt.args.secondList), "intervalIntersection(%v, %v)", tt.args.firstList, tt.args.secondList)
		})
	}
}
