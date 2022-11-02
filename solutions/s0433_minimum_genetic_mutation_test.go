package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_minMutation(t *testing.T) {
	type args struct {
		start string
		end   string
		bank  []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"test1",
			args{
				"AAAAACCC",
				"AACCCCCC",
				[]string{"AAAACCCC", "AAACCCCC", "AACCCCCC"},
			},
			3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, minMutation(tt.args.start, tt.args.end, tt.args.bank), "minMutation(%v, %v, %v)", tt.args.start, tt.args.end, tt.args.bank)
		})
	}
}
