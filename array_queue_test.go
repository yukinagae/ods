package obs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayQueue(t *testing.T) {
	aq := NewArrayQueue()
	assert.Equal(t, 0, aq.n)
	assert.Equal(t, 0, aq.Size())
	assert.Equal(t, 0, aq.j)
	assert.Equal(t, 1, len(aq.a))

	ok := aq.Add("a") // [a]
	assert.True(t, ok)
	assert.Equal(t, ArrayQueue{
		a: []interface{}{"a"},
		j: 0,
		n: 1,
	}, aq)

	ok = aq.Add("b") // [a b]
	assert.True(t, ok)
	assert.Equal(t, ArrayQueue{
		a: []interface{}{"a", "b"},
		j: 0,
		n: 2,
	}, aq)

	x, err := aq.Remove() // [a(removed) b]
	assert.NoError(t, err)
	assert.Equal(t, "a", x)
	assert.Equal(t, ArrayQueue{
		a: []interface{}{"a", "b"},
		j: 1,
		n: 1,
	}, aq)

	x, err = aq.Remove() // []
	assert.NoError(t, err)
	assert.Equal(t, "b", x)
	assert.Equal(t, ArrayQueue{
		a: []interface{}{},
		j: 0,
		n: 0,
	}, aq)
}
