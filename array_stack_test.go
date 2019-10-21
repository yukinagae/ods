package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayStack(t *testing.T) {
	as := NewArrayStack()
	assert.Equal(t, 0, as.n)
	assert.Equal(t, 0, as.Size())
	assert.Equal(t, 1, len(as.a))

	err := as.Add(0, "a") // [a]
	assert.NoError(t, err)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"a"},
		n: 1,
	}, as)
	assert.Equal(t, 1, len(as.a))

	err = as.Add(0, "b") // [b a]
	assert.NoError(t, err)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"b", "a"},
		n: 2,
	}, as)
	assert.Equal(t, 2, len(as.a))

	err = as.Add(1, "c") // [b c a nil]
	assert.NoError(t, err)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"b", "c", "a", nil},
		n: 3,
	}, as)
	assert.Equal(t, 4, len(as.a))

	_, err = as.Remove(0) // [c a nil nil]
	assert.NoError(t, err)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"c", "a", nil, nil},
		n: 2,
	}, as)
	assert.Equal(t, 4, len(as.a))

	err = as.Add(2, "d") // [c a d nil]
	assert.NoError(t, err)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"c", "a", "d", nil},
		n: 3,
	}, as)
	assert.Equal(t, 4, len(as.a))

	value, err := as.Get(0) // [c a d nil] -> c
	assert.NoError(t, err)
	assert.Equal(t, "c", value)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"c", "a", "d", nil},
		n: 3,
	}, as)
	assert.Equal(t, 4, len(as.a))

	value, err = as.Set(0, "foo") // [c a d nil] -> [foo a d nil]
	assert.NoError(t, err)
	assert.Equal(t, "c", value)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"foo", "a", "d", nil},
		n: 3,
	}, as)
	assert.Equal(t, 4, len(as.a))

	value, err = as.Get(0) // [foo a d nil] -> foo
	assert.NoError(t, err)
	assert.Equal(t, "foo", value)
	assert.Equal(t, ArrayStack{
		a: []interface{}{"foo", "a", "d", nil},
		n: 3,
	}, as)
	assert.Equal(t, 4, len(as.a))
}
