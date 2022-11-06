package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_orderlyQueue(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{"cba", 1},
			"acb",
		},
		{
			"test 2",
			args{"baaca", 3},
			"aaabc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, orderlyQueue(tt.args.s, tt.args.k), "orderlyQueue(%v, %v)", tt.args.s, tt.args.k)
		})
	}
}
