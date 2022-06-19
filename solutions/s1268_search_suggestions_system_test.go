package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_suggestedProducts(t *testing.T) {
	type args struct {
		products   []string
		searchWord string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "test1",
			args: args{
				products:   []string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
				searchWord: "mouse",
			},
			want: [][]string{
				{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, suggestedProducts(tt.args.products, tt.args.searchWord), "suggestedProducts(%v, %v)", tt.args.products, tt.args.searchWord)
		})
	}
}
