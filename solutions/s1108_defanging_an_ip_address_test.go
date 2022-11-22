package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_defangIPaddr(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Example 1",
			args: args{
				address: "1.1.1.1",
			},
			want: "1[.]1[.]1[.]1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, defangIPaddr(tt.args.address), "defangIPaddr(%v)", tt.args.address)
		})
	}
}
