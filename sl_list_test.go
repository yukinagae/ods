package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSLList(t *testing.T) {
	list := SLList{
		n: 5,
		head: &Node{
			x: "a",
			next: &Node{
				x: "b",
				next: &Node{
					x: "c",
					next: &Node{
						x: "d",
						next: &Node{
							x:    "e",
							next: nil,
						},
					},
				},
			},
		},
	}
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

	v = list.Pop() // [b c d e x]
	assert.Equal(t, 5, list.n)
	assert.Equal(t, "a", v)
	v, err = list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "b", v)

	v = list.Pop() // [c d e x]
	assert.Equal(t, 4, list.n)
	assert.Equal(t, "b", v)
	v, err = list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "c", v)

	v = list.Push("y") // [y c d e x]
	assert.Equal(t, 5, list.n)
	assert.Equal(t, "y", v)
	v, err = list.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "y", v)
}
