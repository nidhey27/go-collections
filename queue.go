package collections

// Queue defines a generic interface for a queue.
type Queue[T any] interface {
	Enqueue(value T)
	Dequeue() (T, bool)
	Peek() (T, bool)
	Size() int
}

// LinkedListQueue implements Queue using a LinkedList.
type LinkedListQueue[T any] struct {
	list LinkedList[T]
}

// NewLinkedListQueue creates a new LinkedListQueue.
func NewLinkedListQueue[T any]() *LinkedListQueue[T] {
	return &LinkedListQueue[T]{list: LinkedList[T]{}}
}

func (q *LinkedListQueue[T]) Enqueue(value T) {
	q.list.Add(value)
}

func (q *LinkedListQueue[T]) Dequeue() (T, bool) {
	if q.list.Head == nil {
		var zero T
		return zero, false
	}
	value := q.list.Head.Value
	q.list.Head = q.list.Head.Next
	q.list.Size--
	return value, true
}

func (q *LinkedListQueue[T]) Peek() (T, bool) {
	if q.list.Head == nil {
		var zero T
		return zero, false
	}
	return q.list.Head.Value, true
}

func (q *LinkedListQueue[T]) Size() int {
	return q.list.Size
}

// SliceQueue implements Queue using a Slice.
type SliceQueue[T any] struct {
	elements []T
}

// NewSliceQueue creates a new SliceQueue.
func NewSliceQueue[T any]() *SliceQueue[T] {
	return &SliceQueue[T]{elements: []T{}}
}

func (q *SliceQueue[T]) Enqueue(value T) {
	q.elements = append(q.elements, value)
}

func (q *SliceQueue[T]) Dequeue() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false
	}
	value := q.elements[0]
	q.elements = q.elements[1:]
	return value, true
}

func (q *SliceQueue[T]) Peek() (T, bool) {
	if len(q.elements) == 0 {
		var zero T
		return zero, false
	}
	return q.elements[0], true
}

func (q *SliceQueue[T]) Size() int {
	return len(q.elements)
}
