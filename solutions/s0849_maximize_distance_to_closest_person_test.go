package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxDistToClosest(t *testing.T) {
	type args struct {
		seats []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				seats: []int{1, 0, 0, 0, 1, 0, 1},
			},
			want: 2,
		},
		{
			name: "Test 2",
			args: args{
				seats: []int{1, 0, 0, 0},
			},
			want: 3,
		},
		{
			name: "Test 3",
			args: args{
				seats: []int{0, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maxDistToClosest(tt.args.seats), "maxDistToClosest(%v)", tt.args.seats)
		})
	}
}
