package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkIfPangram(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Example 1",
			args: args{
				sentence: "thequickbrownfoxjumpsoverthelazydog",
			},
			want: true,
		},
		{
			name: "Example 2",
			args: args{
				sentence: "leetcode",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, checkIfPangram(tt.args.sentence), "checkIfPangram(%v)", tt.args.sentence)
		})
	}
}
