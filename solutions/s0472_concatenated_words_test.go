package solutions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findAllConcatenatedWordsInADict(t *testing.T) {
	type args struct {
		words []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test 1",
			args{
				[]string{"cat", "cats", "catsdogcats", "dog", "dogcatsdog", "hippopotamuses", "rat", "ratcatdogcat"},
			},
			[]string{"catsdogcats", "dogcatsdog", "ratcatdogcat"},
		},
		{
			"test 2",
			args{
				[]string{"cat", "dog", "catdog"},
			},
			[]string{"catdog"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findAllConcatenatedWordsInADict(tt.args.words), "findAllConcatenatedWordsInADict(%v)", tt.args.words)
		})
	}
}
