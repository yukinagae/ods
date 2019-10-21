package obs

import (
	"errors"
)

// MEMO: BaseList interfaces
// Size() int
// Append(x interface{})
// AddAll(xs []interface{})

func (q ArrayDeque) Size() int {
	return q.n
}

func (q *ArrayDeque) Append(x interface{}) {
	q.Add(q.Size(), x)
}

func (q *ArrayDeque) AddAll(xs []interface{}) {
	for _, x := range xs {
		q.Append(x)
	}
}

// MEMO: ArrayDeque interfaces
// Get(i int) (interface{}, error)
// Set(i int, x interface{}) (interface{}, error)
// Add(i int, x interface{}) error
//Remove(i int) (interface{}, error)

type ArrayDeque struct {
	a []interface{} // init array size is 1
	j int           // index of first element
	n int           // number of elements
}

func NewArrayDeque() ArrayDeque {
	return ArrayDeque{
		a: make([]interface{}, 1),
		j: 0,
		n: 0,
	}
}

func (q ArrayDeque) Get(i int) (interface{}, error) {
	if i < 0 || i >= q.n {
		return nil, errors.New("index error")
	}
	return q.a[(i+q.j)%len(q.a)], nil
}

func (q *ArrayDeque) Set(i int, x interface{}) (interface{}, error) {
	if i < 0 || i >= q.n {
		return nil, errors.New("index error")
	}
	y := q.a[(i+q.j)%len(q.a)]
	q.a[(i+q.j)%len(q.a)] = x
	return y, nil
}

func (q *ArrayDeque) Add(i int, x interface{}) error {
	if i < 0 || i > q.n {
		return errors.New("index error")
	}

	if q.n == len(q.a) {
		q.resize()
	}

	if i < (q.n / 2) {
		if q.j == 0 { // avoid underflow
			q.j = len(q.a) - 1
		} else {
			q.j = (q.j - 1) % len(q.a)
		}
		for k := 0; k < i; k++ {
			q.a[(q.j+k)%len(q.a)] = q.a[(q.j+k+1)%len(q.a)]
		}
	} else {
		for k := q.n; k > i; k-- {
			q.a[(q.j+k)%len(q.a)] = q.a[(q.j+k-1)%len(q.a)]
		}
	}

	q.a[(q.j+i)%len(q.a)] = x
	q.n++
	return nil
}

func (q *ArrayDeque) Remove(i int) (interface{}, error) {
	if i < 0 || i >= q.n {
		return nil, errors.New("index error")
	}

	x := q.a[(q.j+i)%len(q.a)]
	if i < (q.n / 2) {
		for k := i; k > 0; k-- {
			q.a[(q.j+k)%len(q.a)] = q.a[(q.j+k-1)%len(q.a)]
		}
		q.j = (q.j + 1) % len(q.a)
	} else {
		for k := i; k < (q.n - 1); k++ {
			q.a[(q.j+k)%len(q.a)] = q.a[(q.j+k+1)%len(q.a)]
		}
	}

	q.n--

	if len(q.a) >= 3*q.n {
		q.resize()
	}
	return x, nil
}

func (q *ArrayDeque) resize() {
	b := make([]interface{}, 2*q.n)
	for k := 0; k < q.n; k++ {
		b[k] = q.a[(q.j+k)%len(q.a)]
	}
	q.a = b
	q.j = 0
}
