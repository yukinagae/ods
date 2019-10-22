package ods

import (
	"errors"
	"fmt"
)

// MEMO: BaseList interfaces
// Size() int
// Append(x interface{})
// AddAll(xs []interface{})
// Clear()
// AddFirst(x interface{})
// RemoveFirst()
// AddLast(x interface{})
// RemoveLast()
// Insert(i int, x interface{}) error
// Index(x interface{}) int
// RemoceValue(x interface{}) (interface{}, error)

func (l SEList) Size() int {
	return l.n
}

func (l *SEList) Append(x interface{}) {
	l.Add(l.Size(), x)
}

func (l *SEList) AddAll(xs []interface{}) {
	for _, x := range xs {
		l.Append(x)
	}
}

func (l *SEList) Clear() {
	for l.Size() > 0 {
		l.Remove(l.Size() - 1)
	}
}

func (l *SEList) AddFirst(x interface{}) {
	l.Add(0, x)
}

func (l *SEList) RemoveFirst() {
	l.Remove(0)
}

func (l *SEList) AddLast(x interface{}) {
	l.Add(l.Size(), x)
}

func (l *SEList) RemoveLast() {
	l.Remove(l.Size() - 1)
}

func (l *SEList) Insert(i int, x interface{}) error {
	return l.Add(i, x)
}

func (l SEList) Index(x interface{}) (int, error) {
	var u *Node
	for i := 0; i < l.Size(); i++ {
		u = u.next
		if u.x == x {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%v is not in the list", x)
}

func (l *SEList) RemoveValue(x interface{}) (interface{}, error) {
	index, err := l.Index(x)
	if err != nil {
		return nil, err
	}
	return l.Remove(index)
}

// MEMO: SEList interfaces
// Get(i int) (interface{}, error)
// Set(i int, x interface{}) (interface{}, error)
// Add(i int, x interface{}) error
// Remove(i int) (interface{}, error)

type SEList struct {
	n     int
	dummy *SENode
}

type SENode struct {
	d    ArrayDeque
	next *SENode
	prev *SENode
}

func NewSENode(x interface{}) *SENode {
	return &SENode{
		d: ArrayDeque{
			a: make([]interface{}, 10), // TODO: initialized size?
			j: 0,
			n: 0,
		},
		next: nil,
		prev: nil,
	}
}

func NewSEList() SEList {
	dummy := NewSENode(nil)
	dummy.prev = dummy
	dummy.next = dummy

	return SEList{
		n:     0,
		dummy: dummy,
	}
}

// TODO: not yet implemented
func (l SEList) Get(i int) (interface{}, error) {
	if i < 0 || i >= l.n {
		return nil, errors.New("index error")
	}
	return nil, nil
}

// TODO: not yet implemented
func (l *SEList) Set(i int, x interface{}) (interface{}, error) {
	return nil, nil
}

// TODO: not yet implemented
func (l *SEList) Add(i int, x interface{}) error {
	if i < 0 || i > l.n {
		return errors.New("index error")
	}

	return nil
}

// TODO: not yet implemented
func (l *SEList) Remove(i int) (interface{}, error) {
	if i < 0 || i >= l.n {
		return nil, errors.New("index error")
	}

	return nil, nil
}
