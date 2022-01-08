package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseBits(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{
			name: "Example 1",
			args: args{
				num: 43261596,
			},
			want: 964176192,
		},
		{
			name: "Example 2",
			args: args{
				num: 4294967293,
			},
			want: 3221225471,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, reverseBits(tt.args.num), "reverseBits(%v)", tt.args.num)
		})
	}
}
