package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_minPartitions(t *testing.T) {
	type args struct {
		n string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				n: "32",
			},
			want: 3,
		},
		{
			name: "Test 2",
			args: args{
				n: "27346209830709182346",
			},
			want: 9,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minPartitions(tt.args.n), "minPartitions(%v)", tt.args.n)
		})
	}
}
