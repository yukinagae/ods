package ods

import (
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
// RemoceValue(x interface{}) (interface{}, error

func (q DualArrayDeque) Size() int {
	return q.front.Size() + q.back.Size()
}

func (q *DualArrayDeque) Append(x interface{}) {
	q.Add(q.Size(), x)
}

func (q *DualArrayDeque) AddAll(xs []interface{}) {
	for _, x := range xs {
		q.Append(x)
	}
}

func (q *DualArrayDeque) Clear() {
	q.front.Clear()
	q.back.Clear()
}

func (q *DualArrayDeque) AddFirst(x interface{}) {
	q.Add(0, x)
}

func (q *DualArrayDeque) RemoveFirst() {
	q.Remove(0)
}

func (q *DualArrayDeque) AddLast(x interface{}) {
	q.Add(q.Size(), x)
}

func (q *DualArrayDeque) RemoveLast() {
	q.Remove(q.Size() - 1)
}

func (q *DualArrayDeque) Insert(i int, x interface{}) error {
	return q.Add(i, x)
}

func (q DualArrayDeque) Index(x interface{}) (int, error) {

	i, err := q.front.Index(x)
	if err == nil {
		return q.front.Size() - i, nil
	}

	i, err = q.back.Index(x)

	if err == nil {
		return q.front.Size() + i, nil
	}

	return 0, fmt.Errorf("%v is not in the list", x)
}

func (q *DualArrayDeque) RemoveValue(x interface{}) (interface{}, error) {
	i, err := q.Index(x)
	if err != nil {
		return nil, err
	}

	if i < q.front.Size() {
		return q.front.Remove(q.front.Size() - i - 1)
	}
	return q.back.Remove(i - q.front.Size())
}

// MEMO: DualArrayDeque interfaces
// Get(i int) (interface{}, error)
// Set(i int, x interface{}) (interface{}, error)
// Add(i int, x interface{}) error
// Remove(i int) (interface{}, error

type DualArrayDeque struct {
	front ArrayStack
	back  ArrayStack
}

func NewDualArrayDeque() DualArrayDeque {
	return DualArrayDeque{
		front: NewArrayStack(),
		back:  NewArrayStack(),
	}
}

func (q DualArrayDeque) Get(i int) (interface{}, error) {
	if i < q.front.Size() {
		return q.front.Get(q.front.Size() - i - 1)
	}
	return q.back.Get(i - q.front.Size())
}

func (q *DualArrayDeque) Set(i int, x interface{}) (interface{}, error) {
	if i < q.front.Size() {
		return q.front.Set(q.front.Size()-i-1, x)
	}
	return q.back.Set(i-q.front.Size(), x)
}

func (q *DualArrayDeque) Add(i int, x interface{}) error {
	var err error

	if i < q.front.Size() {
		err = q.front.Add(q.front.Size()-i, x)
	} else {
		err = q.back.Add(i-q.front.Size(), x)
	}

	if err != nil {
		return err
	}

	q.balance()
	return nil
}

func (q *DualArrayDeque) Remove(i int) (interface{}, error) {
	var x interface{}
	var err error

	if i < q.front.Size() {
		x, err = q.front.Remove(q.front.Size() - i - 1)

	} else {
		x, err = q.back.Remove(i - q.front.Size())
	}

	if err != nil {
		return nil, err
	}

	q.balance()
	return x, nil
}

func (q *DualArrayDeque) balance() {
	n := q.Size()
	mid := n / 2

	if 3*q.front.Size() < q.back.Size() || 3*q.back.Size() < q.front.Size() {
		f := NewArrayStack()
		for i := 0; i < mid; i++ {
			x, _ := q.Get(mid - i - 1)
			f.Add(i, x)
		}
		b := NewArrayStack()
		for i := 0; i < (n - mid); i++ {
			x, _ := q.Get(mid + i)
			b.Add(i, x)
		}
		q.front = f
		q.back = b
	}
}
