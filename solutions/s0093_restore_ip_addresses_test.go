package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test 1",
			args{"25525511135"},
			[]string{
				"255.255.11.135", "255.255.111.35",
			},
		},
		{
			"test 2",
			args{"0000"},
			[]string{
				"0.0.0.0",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, restoreIpAddresses(tt.args.s), "restoreIpAddresses(%v)", tt.args.s)
		})
	}
}
