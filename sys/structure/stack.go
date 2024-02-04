package structure

// Stack 栈
type Stack[T any] struct {
	content []T
}

func (s *Stack[T]) Push(v T) {
	s.content = append(s.content, v)
}
func (s *Stack[T]) Pop() T {
	l := s.Len()
	v := s.content[l-1]
	s.content = s.content[:l-1]
	return v
}
func (s *Stack[T]) Len() int {
	return len(s.content)
}

func (s *Stack[T]) GetContent() []T {
	v := make([]T, s.Len())
	copy(v, s.content)
	return v
}
