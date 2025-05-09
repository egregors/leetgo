package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minTimeToReach3342(t *testing.T) {
	type args struct {
		moveTime [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test Case 1",
			args: args{
				moveTime: [][]int{
					{0, 4},
					{4, 4},
				},
			},
			want: 7,
		},
		{
			name: "Test Case 2",
			args: args{
				moveTime: [][]int{
					{0, 0, 0, 0},
					{0, 0, 0, 0},
				},
			},
			want: 6,
		},
		{
			name: "Test Case 3",
			args: args{
				moveTime: [][]int{
					{0, 1},
					{1, 2},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minTimeToReach3342(tt.args.moveTime), "minTimeToReach3342(%v)", tt.args.moveTime)
		})
	}
}
