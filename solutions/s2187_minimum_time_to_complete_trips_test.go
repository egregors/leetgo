package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minimumTime(t *testing.T) {
	type args struct {
		time       []int
		totalTrips int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{
			"test 1",
			args{
				[]int{1, 2, 3},
				5,
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minimumTime(tt.args.time, tt.args.totalTrips), "minimumTime(%v, %v)", tt.args.time, tt.args.totalTrips)
		})
	}
}
