package collections

// Stack defines a generic interface for a stack.
type Stack[T any] interface {
	Push(value T)
	Pop() (T, bool)
	Peek() (T, bool)
	Size() int
}

// LinkedListStack implements Stack using a LinkedList.
type LinkedListStack[T any] struct {
	list LinkedList[T]
}

// NewLinkedListStack creates a new LinkedListStack.
func NewLinkedListStack[T any]() *LinkedListStack[T] {
	return &LinkedListStack[T]{list: LinkedList[T]{}}
}

func (s *LinkedListStack[T]) Push(value T) {
	newNode := &Node[T]{Value: value, Next: s.list.Head}
	s.list.Head = newNode
	s.list.Size++
}

func (s *LinkedListStack[T]) Pop() (T, bool) {
	if s.list.Head == nil {
		var zero T
		return zero, false
	}
	value := s.list.Head.Value
	s.list.Head = s.list.Head.Next
	s.list.Size--
	return value, true
}

func (s *LinkedListStack[T]) Peek() (T, bool) {
	if s.list.Head == nil {
		var zero T
		return zero, false
	}
	return s.list.Head.Value, true
}

func (s *LinkedListStack[T]) Size() int {
	return s.list.Size
}

// SliceStack implements Stack using a Slice.
type SliceStack[T any] struct {
	elements []T
}

// NewSliceStack creates a new SliceStack.
func NewSliceStack[T any]() *SliceStack[T] {
	return &SliceStack[T]{elements: []T{}}
}

func (s *SliceStack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}

func (s *SliceStack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	value := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return value, true
}

func (s *SliceStack[T]) Peek() (T, bool) {
	if len(s.elements) == 0 {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *SliceStack[T]) Size() int {
	return len(s.elements)
}
