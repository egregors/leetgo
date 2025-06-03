package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_toGoatLatin(t *testing.T) {
	type args struct {
		sentence string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		//Input: sentence = "I speak Goat Latin"
		//Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
		{
			name: "Test1",
			args: args{sentence: "I speak Goat Latin"},
			want: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa",
		},
		//Input: sentence = "The quick brown fox jumped over the lazy dog"
		//Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
		{
			name: "Test2",
			args: args{sentence: "The quick brown fox jumped over the lazy dog"},
			want: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, toGoatLatin(tt.args.sentence), "toGoatLatin(%v)", tt.args.sentence)
		})
	}
}
