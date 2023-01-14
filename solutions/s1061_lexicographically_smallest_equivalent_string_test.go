package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_smallestEquivalentString(t *testing.T) {
	type args struct {
		s1      string
		s2      string
		baseStr string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"test 1",
			args{
				"parker",
				"morris",
				"parser",
			},
			"makkek",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, smallestEquivalentString(tt.args.s1, tt.args.s2, tt.args.baseStr), "smallestEquivalentString(%v, %v, %v)", tt.args.s1, tt.args.s2, tt.args.baseStr)
		})
	}
}
