package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_matchPlayersAndTrainers(t *testing.T) {
	type args struct {
		players  []int
		trainers []int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		//Input: players = [4,7,9], trainers = [8,2,5,8]
		//Output: 2
		{
			name: "Test1",
			args: args{
				players:  []int{4, 7, 9},
				trainers: []int{8, 2, 5, 8},
			},
			wantAns: 2,
		},
		//Input: players = [1,1,1], trainers = [10]
		//Output: 1
		{
			name: "Test2",
			args: args{
				players:  []int{1, 1, 1},
				trainers: []int{10},
			},
			wantAns: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.wantAns, matchPlayersAndTrainers(tt.args.players, tt.args.trainers), "matchPlayersAndTrainers(%v, %v)", tt.args.players, tt.args.trainers)
		})
	}
}
