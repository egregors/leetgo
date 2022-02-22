package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_titleToNumber(t *testing.T) {
	type args struct {
		columnTitle string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test 1",
			args: args{
				columnTitle: "A",
			},
			want: 1,
		},
		{
			name: "Test 2",
			args: args{
				columnTitle: "AB",
			},
			want: 28,
		},
		{
			name: "Test 3",
			args: args{
				columnTitle: "ZY",
			},
			want: 701,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, titleToNumber(tt.args.columnTitle), "titleToNumber(%v)", tt.args.columnTitle)
		})
	}
}
