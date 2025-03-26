package stack

type Stack[T any] struct {
	stack []T
}

func New[T any]() *Stack[T] {
	return new(Stack[T])
}

func (s *Stack[T]) Push(value T) {
	s.stack = append(s.stack, value)
}

func (s *Stack[T]) Pop() T {
	last := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]

	return last
}

func (s *Stack[T]) Len() int {
	return len(s.stack)
}
