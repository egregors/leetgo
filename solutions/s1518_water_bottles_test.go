package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numWaterBottles(t *testing.T) {
	type args struct {
		numBottles  int
		numExchange int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test1",
			args: args{
				numBottles:  9,
				numExchange: 3,
			},
			want: 13,
		},
		{
			name: "test2",
			args: args{
				numBottles:  15,
				numExchange: 4,
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, numWaterBottles(tt.args.numBottles, tt.args.numExchange), "numWaterBottles(%v, %v)", tt.args.numBottles, tt.args.numExchange)
		})
	}
}
