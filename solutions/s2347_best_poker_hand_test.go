package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_bestHand(t *testing.T) {
	type args struct {
		ranks []int
		suits []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//Input: ranks = [13,2,3,1,9], suits = ["a","a","a","a","a"]
		//Output: "Flush"
		//Input: ranks = [4,4,2,4,4], suits = ["d","a","a","b","c"]
		//Output: "Three of a Kind"
		//Input: ranks = [10,10,2,12,9], suits = ["a","b","c","a","d"]
		//Output: "Pair"
		{
			name: "Test1",
			args: args{
				ranks: []int{13, 2, 3, 1, 9},
				suits: []byte{'a', 'a', 'a', 'a', 'a'},
			},
			want: "Flush",
		},
		{
			name: "Test2",
			args: args{
				ranks: []int{4, 4, 2, 4, 4},
				suits: []byte{'d', 'a', 'a', 'b', 'c'},
			},
			want: "Three of a Kind",
		},
		{
			name: "Test3",
			args: args{
				ranks: []int{10, 10, 2, 12, 9},
				suits: []byte{'a', 'b', 'c', 'a', 'd'},
			},
			want: "Pair",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, bestHand(tt.args.ranks, tt.args.suits), "bestHand(%v, %v)", tt.args.ranks, tt.args.suits)
		})
	}
}
