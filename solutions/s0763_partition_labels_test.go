package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_partitionLabels(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test 1",
			args: args{
				s: "ababcbacadefegdehijhklij",
			},
			want: []int{9, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, partitionLabels(tt.args.s), "partitionLabels(%v)", tt.args.s)
		})
	}
}
