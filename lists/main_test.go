package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewList(t *testing.T) {
	assert := assert.New(t)
	list := New()
	node := list.(*List)
	assert.Equal(node.node, "Head")
	assert.Nil(node.link)
}
func TestAppendList(t *testing.T) {
	assert := assert.New(t)
	list := New()
	list.append(12)
	assert.Equal(2, list.length())

	list.append("hello world")
	list.append("12")
	list.append(123)

	assert.Equal(5, list.length())
}
