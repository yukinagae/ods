package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootishArrayStack(t *testing.T) {
	ras := NewRootishArrayStack()
	assert.Equal(t, 0, ras.n)
	assert.Equal(t, 0, ras.Size())
	assert.Equal(t, 0, ras.blocks.Size())
	assert.Equal(t, 1, len(ras.blocks.a))

	// TODO: more tests
}
