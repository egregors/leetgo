package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec(t *testing.T) {
	root := &TreeNode{
		1,
		&TreeNode{
			2,
			nil,
			nil,
		},
		&TreeNode{
			3,
			&TreeNode{
				4,
				nil,
				nil,
			},
			&TreeNode{
				5,
				nil,
				nil,
			},
		},
	}

	codec := NewCodec()
	data := codec.serialize(root)
	assert.Equal(t, data, "[1,2,3,null,null,4,5]")
	assert.Equal(t, root, codec.deserialize(data))
	assert.Equal(t, root, codec.deserialize(codec.serialize(root)))
}
