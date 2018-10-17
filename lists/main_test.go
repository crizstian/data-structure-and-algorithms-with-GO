package lists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppendList(t *testing.T) {
	assert := assert.New(t)
	list := New()
	list.append(12)
	list.append("hello world")
	list.append("what")
	list.append(123)
	assert.Equal(4, list.len)
	list.display()
}
func TestRemoveElementList(t *testing.T) {
	assert := assert.New(t)
	list := New()
	list.append(12)
	list.append("hello world")
	list.append("what")
	list.append(123)
	list.remove("hello world")
	list.remove("what")
	assert.Equal(2, list.len)
	list.remove(12)
	assert.Equal(1, list.len)
	list.display()
}
