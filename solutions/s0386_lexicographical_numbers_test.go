package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lexicalOrder(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		//Input: n = 13
		//Output: [1,10,11,12,13,2,3,4,5,6,7,8,9]
		{
			name: "Test 1",
			args: args{n: 13},
			want: []int{1, 10, 11, 12, 13, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		//Input: n = 2
		//Output: [1,2]
		{
			name: "Test 2",
			args: args{n: 2},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, lexicalOrder(tt.args.n), "lexicalOrder(%v)", tt.args.n)
		})
	}
}
