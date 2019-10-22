package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDualArrayDeque(t *testing.T) {
	daq := NewDualArrayDeque()
	assert.Equal(t, 0, daq.front.Size())
	assert.Equal(t, 0, daq.back.Size())
	assert.Equal(t, 1, len(daq.front.a))
	assert.Equal(t, 1, len(daq.back.a))

	daq = DualArrayDeque{
		front: ArrayStack{
			a: []interface{}{"b", "a", nil, nil, nil},
			n: 2,
		},
		back: ArrayStack{
			a: []interface{}{"c", "d", nil, nil, nil},
			n: 2,
		},
	}

	v, err := daq.Get(0)
	assert.NoError(t, err)
	assert.Equal(t, "a", v)
	v, err = daq.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, "b", v)
	v, err = daq.Get(2)
	assert.NoError(t, err)
	assert.Equal(t, "c", v)
	v, err = daq.Get(3)
	assert.NoError(t, err)
	assert.Equal(t, "d", v)

	err = daq.Add(3, "x")
	assert.NoError(t, err)
	assert.Equal(t, DualArrayDeque{
		front: ArrayStack{
			a: []interface{}{"b", "a", nil, nil, nil},
			n: 2,
		},
		back: ArrayStack{
			a: []interface{}{"c", "x", "d", nil, nil},
			n: 3,
		},
	}, daq)

	err = daq.Add(4, "y")
	assert.NoError(t, err)
	assert.Equal(t, DualArrayDeque{
		front: ArrayStack{
			a: []interface{}{"b", "a", nil, nil, nil},
			n: 2,
		},
		back: ArrayStack{
			a: []interface{}{"c", "x", "y", "d", nil},
			n: 4,
		},
	}, daq)

	v, err = daq.Remove(0)
	assert.Equal(t, "a", v)
	assert.NoError(t, err)
	assert.Equal(t, DualArrayDeque{
		front: ArrayStack{
			a: []interface{}{"c", "b"},
			n: 2,
		},
		back: ArrayStack{
			a: []interface{}{"x", "y", "d", nil},
			n: 3,
		},
	}, daq)

}
