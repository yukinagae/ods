package ods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSEList(t *testing.T) {

	list := NewSEList()
	assert.Equal(t, 0, list.n) // []

	// TODO: more tests

}
