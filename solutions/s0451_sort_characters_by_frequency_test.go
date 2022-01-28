package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test 1",
			args: args{
				s: "too",
			},
			want: "oot",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, frequencySort(tt.args.s), "frequencySort(%v)", tt.args.s)
		})
	}
}
