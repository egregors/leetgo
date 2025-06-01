package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_distributeCandies(t *testing.T) {
	type args struct {
		n     int
		limit int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		//Input: n = 5, limit = 2
		//Output: 3
		{
			name: "Test 1",
			args: args{
				n:     5,
				limit: 2,
			},
			want: 3,
		},
		//Input: n = 3, limit = 3
		//Output: 10
		{
			name: "Test 2",
			args: args{
				n:     3,
				limit: 3,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, distributeCandies(tt.args.n, tt.args.limit), "distributeCandies(%v, %v)", tt.args.n, tt.args.limit)
		})
	}
}
