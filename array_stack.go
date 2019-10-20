package main

import (
	"errors"
	"fmt"
)

// TODO: BaseList
// Size() uint
// Append(x interface{})
// AddAll(xs []interface{})
// Clear()
// AddFirst(x interface{})
// RemoveFirst()
// AddLast(x interface{})
// RemoveLast()
// Insert(i uint, x interface{}) error
// Index(x interface{}) uint
// RemoceValue(x interface{}) (interface{}, error)

func (s ArrayStack) Size() uint {
	return s.n
}

func (s *ArrayStack) Append(x interface{}) {
	s.Add(s.Size(), x)
}

func (s *ArrayStack) AddAll(xs []interface{}) {
	for _, x := range xs {
		s.Append(x)
	}
}

func (s *ArrayStack) Clear() {
	for s.Size() > 0 {
		s.Remove(s.Size() - 1)
	}
}

func (s *ArrayStack) AddFirst(x interface{}) {
	s.Add(0, x)
}

func (s *ArrayStack) RemoveFirst() {
	s.Remove(0)
}

func (s *ArrayStack) AddLast(x interface{}) {
	s.Add(s.Size(), x)
}

func (s *ArrayStack) RemoveLast() {
	s.Remove(s.Size() - 1)
}

func (s *ArrayStack) Insert(i uint, x interface{}) error {
	return s.Add(i, x)
}

func (s ArrayStack) Index(x interface{}) (uint, error) {
	for index, element := range s.a {
		if element == x {
			return uint(index), nil
		}
	}
	return 0, fmt.Errorf("%v is not in the list", x)
}

func (s *ArrayStack) RemoceValue(x interface{}) (interface{}, error) {
	index, err := s.Index(x)
	if err != nil {
		return nil, err
	}
	return s.Remove(index)
}

// ArrayStack implementation
type ArrayStack struct {
	a []interface{} // init array size is 1
	n uint          // number of elements
}

func NewArrayStack() ArrayStack {
	return ArrayStack{
		a: make([]interface{}, 1),
		n: 0,
	}
}

func (s ArrayStack) Get(i uint) (interface{}, error) {
	if i >= s.n {
		return nil, errors.New("index error")
	}
	return s.a[i], nil
}

func (s *ArrayStack) Set(i uint, x interface{}) (interface{}, error) {
	if i >= s.n {
		return nil, errors.New("index error")
	}
	y := s.a[i]
	s.a[i] = x
	return y, nil
}

func (s *ArrayStack) Add(i uint, x interface{}) error {
	if i > s.n {
		return errors.New("index error")
	}

	if s.n == uint(len(s.a)) {
		s.resize()
	}

	// see: https://github.com/golang/go/wiki/SliceTricks
	// see: https://blog.golang.org/go-slices-usage-and-internals
	copy(s.a[i+1:s.n+1], s.a[i:s.n])

	s.a[i] = x
	s.n++
	return nil
}

func (s *ArrayStack) Remove(i uint) (interface{}, error) {
	if i >= s.n {
		return nil, errors.New("index error")
	}
	x := s.a[i]
	s.a = append(s.a[:i], s.a[i+1:]...)
	s.a = append(s.a, nil) // fill nil to the end
	s.n--
	if uint(len(s.a)) >= 3*s.n {
		s.resize()
	}
	return x, nil
}

func (s *ArrayStack) resize() {
	b := make([]interface{}, 2*s.n)
	copy(b[:s.n], s.a[:s.n])
	s.a = b
}

func main() {
	as := NewArrayStack()
	fmt.Println(as.n)
	fmt.Println(as.Size())
	fmt.Println(len(as.a))

	err := as.Add(0, "a") // [a]
	if err != nil {
		panic(err)
	}
	fmt.Println(as)
	err = as.Add(0, "b") // [b a]
	if err != nil {
		panic(err)
	}
	fmt.Println(as)
	err = as.Add(1, "c") // [b c a nil]
	if err != nil {
		panic(err)
	}
	fmt.Println(as)
	_, err = as.Remove(0) // [c a nil nil]
	if err != nil {
		panic(err)
	}
	fmt.Println(as)
	err = as.Add(2, "d") // [c a d nil]
	if err != nil {
		panic(err)
	}
	fmt.Println(as)

	value, _ := as.Get(0)
	fmt.Println(value)

	value, _ = as.Set(0, "foo")
	fmt.Println(value)

	value, _ = as.Get(0)
	fmt.Println(value)

	fmt.Println(as)
}
