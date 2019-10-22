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

func (l SLList) Size() int {
	return l.n
}

func (l *SLList) Append(x interface{}) {
	l.Add(l.Size(), x)
}

func (l *SLList) AddAll(xs []interface{}) {
	for _, x := range xs {
		l.Append(x)
	}
}

func (l *SLList) Clear() {
	for l.Size() > 0 {
		l.Remove(l.Size() - 1)
	}
}

func (l *SLList) AddFirst(x interface{}) {
	l.Add(0, x)
}

func (l *SLList) RemoveFirst() {
	l.Remove(0)
}

func (l *SLList) AddLast(x interface{}) {
	l.Add(l.Size(), x)
}

func (l *SLList) RemoveLast() {
	l.Remove(l.Size() - 1)
}

func (l *SLList) Insert(i int, x interface{}) error {
	return l.Add(i, x)
}

func (l SLList) Index(x interface{}) (int, error) {
	var u *Node
	for i := 0; i < l.Size(); i++ {
		u = u.next
		if u.x == x {
			return i, nil
		}
	}
	return 0, fmt.Errorf("%v is not in the list", x)
}

func (l *SLList) RemoveValue(x interface{}) (interface{}, error) {
	index, err := l.Index(x)
	if err != nil {
		return nil, err
	}
	return l.Remove(index)
}

// MEMO: SLList interfaces
// Get(i int) (interface{}, error)
// Set(i int, x interface{}) (interface{}, error)
// Add(i int, x interface{}) error
// Remove(i int) (interface{}, error)

type SLList struct {
	n    int
	head *Node
	tail *Node
}

type Node struct {
	x    interface{}
	next *Node
}

func NewNode(x interface{}) *Node {
	return &Node{
		x:    x,
		next: nil,
	}
}

func NewSLList() SLList {
	return SLList{
		n:    0,
		head: NewNode(nil),
		tail: NewNode(nil),
	}
}

func (s SLList) Get(i int) (interface{}, error) {
	if i < 0 || i >= s.n {
		return nil, errors.New("index error")
	}
	return s.getNode(i).x, nil
}

func (s *SLList) Set(i int, x interface{}) (interface{}, error) {
	if i < 0 || i >= s.n {
		return nil, errors.New("index error")
	}
	u := s.getNode(i)
	y := u.x
	u.x = x
	return y, nil
}

func (s *SLList) Add(i int, x interface{}) error {
	if i < 0 || i > s.n {
		return errors.New("index error")
	}

	if i == 0 {
		s.Push(x)
		return nil
	}

	u := s.head

	for k := 0; k < (i - 1); k++ {
		u = u.next
	}

	w := NewNode(x)
	w.next = u.next
	u.next = w
	s.n++
	return nil
}

func (s *SLList) Remove(i int) (interface{}, error) {
	if i < 0 || i >= s.n {
		return nil, errors.New("index error")
	}

	if i == 0 {
		x := s.Pop()
		if x == nil {
			return nil, errors.New("no element to be removed")
		}
		return x, nil
	}

	u := s.head

	for k := 0; k < (i - 1); k++ {
		u = u.next
	}

	w := u.next
	u.next = u.next.next

	s.n--

	if s.n == 0 {
		s.tail = nil
	}

	return w.x, nil
}

func (s SLList) getNode(i int) *Node {
	u := s.head
	for k := 0; k < i; k++ {
		u = u.next
	}
	return u
}

func (s *SLList) Push(x interface{}) interface{} {
	u := NewNode(x)

	u.next = s.head
	s.head = u

	if s.n == 0 {
		s.tail = u
	}

	s.n++
	return x
}

func (s *SLList) Pop() interface{} {
	if s.n == 0 {
		return nil
	}

	x := s.head.x
	s.head = s.head.next

	s.n--

	if s.n == 0 {
		s.tail = nil
	}

	return x
}
