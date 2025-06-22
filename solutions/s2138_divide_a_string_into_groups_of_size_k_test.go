package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_divideString(t *testing.T) {
	type args struct {
		s    string
		k    int
		fill byte
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		//Input: s = "abcdefghi", k = 3, fill = "x"
		//Output: ["abc","def","ghi"]
		{
			name: "Test 1",
			args: args{
				s:    "abcdefghi",
				k:    3,
				fill: 'x',
			},
			want: []string{"abc", "def", "ghi"},
		},
		//Input: s = "abcdefghij", k = 3, fill = "x"
		//Output: ["abc","def","ghi","jxx"]
		{
			name: "Test 2",
			args: args{
				s:    "abcdefghij",
				k:    3,
				fill: 'x',
			},
			want: []string{"abc", "def", "ghi", "jxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, divideString(tt.args.s, tt.args.k, tt.args.fill), "divideString(%v, %v, %v)", tt.args.s, tt.args.k, tt.args.fill)
		})
	}
}
