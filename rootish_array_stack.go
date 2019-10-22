package ods

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

func (s RootishArrayStack) Size() int {
	return s.n
}

func (s *RootishArrayStack) Append(x interface{}) {
	s.Add(s.Size(), x)
}

func (s *RootishArrayStack) AddAll(xs []interface{}) {
	for _, x := range xs {
		s.Append(x)
	}
}

func (s *RootishArrayStack) Clear() {
	s.blocks.Clear()
	s.n = 0
}

func (s *RootishArrayStack) AddFirst(x interface{}) {
	s.Add(0, x)
}

func (s *RootishArrayStack) RemoveFirst() {
	s.Remove(0)
}

func (s *RootishArrayStack) AddLast(x interface{}) {
	s.Add(s.Size(), x)
}

func (s *RootishArrayStack) RemoveLast() {
	s.Remove(s.Size() - 1)
}

func (s *RootishArrayStack) Insert(i int, x interface{}) error {
	return s.Add(i, x)
}

// TODO: not yet implemented
func (s RootishArrayStack) Index(x interface{}) (int, error) {
	return 0, nil
}

// TODO: not yet implemented
func (s *RootishArrayStack) RemoveValue(x interface{}) (interface{}, error) {
	return nil, nil
}

// MEMO: RootishArrayStack interfaces
// Get(i int) (interface{}, error)
// Set(i int, x interface{}) (interface{}, error)
// Add(i int, x interface{}) error
// Remove(i int) (interface{}, error)

type RootishArrayStack struct {
	n      int
	blocks ArrayStack
}

func NewRootishArrayStack() RootishArrayStack {
	return RootishArrayStack{
		n:      0,
		blocks: NewArrayStack(),
	}
}

// TODO: not yet implemented
func (s RootishArrayStack) Get(i int) (interface{}, error) {
	return nil, nil
}

// TODO: not yet implemented
func (s *RootishArrayStack) Set(i int, x interface{}) (interface{}, error) {
	return nil, nil
}

// TODO: not yet implemented
func (s *RootishArrayStack) Add(i int, x interface{}) error {
	return nil
}

// TODO: not yet implemented
func (s RootishArrayStack) Remove(i int) (interface{}, error) {
	return nil, nil
}

// TODO: not yet implemented
func (s RootishArrayStack) i2b(i int) int {
	return 0
}

// TODO: not yet implemented
func (s *RootishArrayStack) grow(i int) {
	return 0
}

// TODO: not yet implemented
func (s *RootishArrayStack) shrink(i int) {
}
