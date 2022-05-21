package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_coinChange(t *testing.T) {
	type args struct {
		coins  []int
		amount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				coins:  []int{1, 2, 5},
				amount: 11,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, coinChange(tt.args.coins, tt.args.amount), "coinChange(%v, %v)", tt.args.coins, tt.args.amount)
		})
	}
}
