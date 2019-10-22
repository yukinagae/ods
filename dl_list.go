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

func (l DLList) Size() int {
	return l.n
}

func (l *DLList) Append(x interface{}) {
	l.Add(l.Size(), x)
}

func (l *DLList) AddAll(xs []interface{}) {
	for _, x := range xs {
		l.Append(x)
	}
}

func (l *DLList) Clear() {
	for l.Size() > 0 {
		l.Remove(l.Size() - 1)
	}
}

func (l *DLList) AddFirst(x interface{}) {
	l.Add(0, x)
}

func (l *DLList) RemoveFirst() {
	l.Remove(0)
}

func (l *DLList) AddLast(x interface{}) {
	l.Add(l.Size(), x)
}

func (l *DLList) RemoveLast() {
	l.Remove(l.Size() - 1)
}

func (l *DLList) Insert(i int, x interface{}) error {
	return l.Add(i, x)
}

func (l DLList) Index(x interface{}) (int, error) {
	var u *Node
	for i := 0; i < l.Size(); i++ {
		u = u.next
		if u.x == x {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%v is not in the list", x)
}

func (l *DLList) RemoveValue(x interface{}) (interface{}, error) {
	index, err := l.Index(x)
	if err != nil {
		return nil, err
	}
	return l.Remove(index)
}

// MEMO: DLList interfaces
// Get(i int) (interface{}, error)
// Set(i int, x interface{}) (interface{}, error)
// Add(i int, x interface{}) error
// Remove(i int) (interface{}, error)

type DLList struct {
	n     int
	dummy *DLNode
}

type DLNode struct {
	x    interface{}
	next *DLNode
	prev *DLNode
}

func NewDLNode(x interface{}) *DLNode {
	return &DLNode{
		x:    x,
		next: nil,
		prev: nil,
	}
}

func NewDLList() DLList {
	dummy := NewDLNode(nil)
	dummy.prev = dummy
	dummy.next = dummy

	return DLList{
		n:     0,
		dummy: dummy,
	}
}

func (l DLList) Get(i int) (interface{}, error) {
	if i < 0 || i >= l.n {
		return nil, errors.New("index error")
	}
	return l.getNode(i).x, nil
}

func (l *DLList) Set(i int, x interface{}) (interface{}, error) {
	if i < 0 || i >= l.n {
		return nil, errors.New("index error")
	}
	u := l.getNode(i)
	y := u.x
	u.x = x
	return y, nil
}

func (l *DLList) Add(i int, x interface{}) error {
	if i < 0 || i > l.n {
		return errors.New("index error")
	}

	w := l.getNode(i)
	u := NewDLNode(x)

	// [p -> (w), (p) <- w]
	u.prev = w.prev // [p -> (w), (p) <- u, (p) <- w]
	u.next = w      // [p -> (w), (p) <- u -> (w), (p) <- w]
	u.next.prev = u // [p -> (w), (p) <- u -> (w), (u) <- w]
	u.prev.next = u // [p -> (u), (p) <- u -> (w), (u) <- w]
	l.n++

	return nil
}

func (l *DLList) Remove(i int) (interface{}, error) {
	if i < 0 || i >= l.n {
		return nil, errors.New("index error")
	}

	w := l.getNode(i)
	w.prev.next = w.next
	w.next.prev = w.prev
	l.n--

	return w.x, nil
}

func (l DLList) getNode(i int) *DLNode {
	if i < (l.n / 2) {
		p := l.dummy.next
		for k := 0; k < i; k++ {
			p = p.next
		}
		return p
	}

	p := l.dummy
	for k := l.n; k > i; k-- {
		p = p.prev
	}
	return p
}
