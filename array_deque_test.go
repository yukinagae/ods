package obs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArrayDeque(t *testing.T) {
	ad := ArrayDeque{
		a: []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", nil, nil, nil, nil},
		j: 0,
		n: 8,
	}
	assert.Equal(t, ArrayDeque{
		a: []interface{}{"a", "b", "c", "d", "e", "f", "g", "h", nil, nil, nil, nil},
		j: 0,
		n: 8,
	}, ad)

	value, err := ad.Remove(2) // [a(garbage) a b d e f g h nil nil nil nil]
	assert.NoError(t, err)
	assert.Equal(t, "c", value)
	assert.Equal(t, ArrayDeque{
		a: []interface{}{"a", "a", "b", "d", "e", "f", "g", "h", nil, nil, nil, nil},
		j: 1,
		n: 7,
	}, ad)

	err = ad.Add(4, "x") // [a(garbage) a b d e x f g h nil nil nil]
	assert.NoError(t, err)
	assert.Equal(t, ArrayDeque{
		a: []interface{}{"a", "a", "b", "d", "e", "x", "f", "g", "h", nil, nil, nil},
		j: 1,
		n: 8,
	}, ad)

	err = ad.Add(3, "y") // [a b d y e x f g h nil nil nil]
	assert.NoError(t, err)
	assert.Equal(t, ArrayDeque{
		a: []interface{}{"a", "b", "d", "y", "e", "x", "f", "g", "h", nil, nil, nil},
		j: 0,
		n: 9,
	}, ad)

	err = ad.Add(3, "z") // [b d z y e x f g h nil nil a]
	assert.NoError(t, err)
	assert.Equal(t, ArrayDeque{
		a: []interface{}{"b", "d", "z", "y", "e", "x", "f", "g", "h", nil, nil, "a"},
		j: 11,
		n: 10,
	}, ad)
}
