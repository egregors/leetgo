package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findAllRecipes(t *testing.T) {
	type args struct {
		recipes     []string
		ingredients [][]string
		supplies    []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"test 1",
			args{
				[]string{"bread", "sandwich"},
				[][]string{{"yeast", "flour"}, {"bread", "meat"}},
				[]string{"yeast", "flour", "meat"},
			},
			[]string{"bread", "sandwich"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, findAllRecipes(tt.args.recipes, tt.args.ingredients, tt.args.supplies), "findAllRecipes(%v, %v, %v)", tt.args.recipes, tt.args.ingredients, tt.args.supplies)
		})
	}
}
