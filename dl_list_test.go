package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDLList(t *testing.T) {

	list := NewDLList()
	assert.Equal(t, 0, list.n) // []

	list.Append("a")
	list.Append("b")
	list.Append("c")
	list.Append("d")
	list.Append("e")
	assert.Equal(t, 5, list.n) // [a b c d e]

	v, err := list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "a", v)
	v, err = list.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, "b", v)
	v, err = list.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, "c", v)
	v, err = list.Get(3)
	assert.NoError(t, err)
	assert.Equal(t, "d", v)
	v, err = list.Get(4)
	assert.NoError(t, err)
	assert.Equal(t, "e", v)
	_, err = list.Get(5)
	assert.Error(t, err)

	list.Append("x") // [a b c d e x]
	assert.Equal(t, 6, list.n)
	v, err = list.Get(5)
	assert.NoError(t, err)
	assert.Equal(t, "x", v)

	list.RemoveFirst() // [b c d e x]
	assert.Equal(t, 5, list.n)
	v, err = list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "b", v)

	list.RemoveFirst() // [c d e x]
	assert.Equal(t, 4, list.n)
	v, err = list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "c", v)

	list.AddFirst("y") // [y c d e x]
	assert.Equal(t, 5, list.n)
	v, err = list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "y", v)
}
