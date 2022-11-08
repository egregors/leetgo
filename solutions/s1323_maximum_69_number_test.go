package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maximum69Number(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Example 1",
			args: args{
				num: 9669,
			},
			want: 9969,
		},
		{
			name: "Example 2",
			args: args{
				num: 9996,
			},
			want: 9999,
		},
		{
			name: "Example 3",
			args: args{
				num: 9999,
			},
			want: 9999,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, maximum69Number(tt.args.num), "maximum69Number(%v)", tt.args.num)
		})
	}
}
