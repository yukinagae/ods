package obs

import (
	"errors"
)

// MEMO: BaseSet interfaces
// Size() int
// AddAll(xs []interface{})

func (q ArrayQueue) Size() int {
	return q.n
}

func (q *ArrayQueue) AddAll(xs []interface{}) {
	for _, x := range xs {
		q.Add(x)
	}
}

// MEMO: ArrayQueue interfaces
// Add(x interface{}) bool
//Remove() (interface{}, error)

type ArrayQueue struct {
	a []interface{} // init array size is 1
	j int           // index of first element
	n int           // number of elements
}

func NewArrayQueue() ArrayQueue {
	return ArrayQueue{
		a: make([]interface{}, 1),
		j: 0,
		n: 0,
	}
}

func (q *ArrayQueue) Add(x interface{}) bool {
	if q.n+1 > len(q.a) {
		q.resize()
	}
	q.a[(q.j+q.n)%len(q.a)] = x
	q.n++
	return true
}

func (q *ArrayQueue) Remove() (interface{}, error) {
	if q.n == 0 {
		return nil, errors.New("index error")
	}
	x := q.a[q.j]
	q.j = (q.j + 1) % len(q.a)
	q.n--

	if len(q.a) >= 3*q.n {
		q.resize()
	}
	return x, nil
}

func (q *ArrayQueue) resize() {
	b := make([]interface{}, 2*q.n)
	for k := 0; k < q.n; k++ {
		b[k] = q.a[(q.j+k)%len(q.a)]
	}
	q.a = b
	q.j = 0
}
