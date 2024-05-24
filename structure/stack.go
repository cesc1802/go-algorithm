package structure

type Stack[T any] struct {
	values []T
}

func New[T any]() *Stack[T] {
	return &Stack[T]{
		values: make([]T, 0),
	}
}

func (s *Stack[T]) IsEmpty() bool { return len(s.values) == 0 }

func (s *Stack[T]) Push(value T) {
	s.values = append(s.values, value)
}

func (s *Stack[T]) Pop() (value T) {
	if s.IsEmpty() {
		return
	}

	size := len(s.values)

	//Get last value in stack
	value = s.values[size-1]
	//remove last value out stack
	s.values = s.values[:size-1]

	return
}
