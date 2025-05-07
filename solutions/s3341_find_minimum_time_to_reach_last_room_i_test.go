package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minTimeToReach(t *testing.T) {
	type args struct {
		moveTime [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				moveTime: [][]int{
					{0, 4},
					{4, 4},
				},
			},
			want: 6,
		},
		{
			name: "Test 2",
			args: args{
				moveTime: [][]int{
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				moveTime: [][]int{
					{0, 1},
					{1, 2},
				},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minTimeToReach(tt.args.moveTime), "minTimeToReach(%v)", tt.args.moveTime)
		})
	}
}
