package collections

import "testing"

func TestLinkedListQueue(t *testing.T) {
	queue := NewLinkedListQueue[int]()
	queue.Enqueue(1)
	queue.Enqueue(2)

	if queue.Size() != 2 {
		t.Errorf("Expected size 2, got %d", queue.Size())
	}

	value, ok := queue.Dequeue()
	if !ok || value != 1 {
		t.Errorf("Expected to dequeue 1, got %v", value)
	}

	value, ok = queue.Peek()
	if !ok || value != 2 {
		t.Errorf("Expected to peek 2, got %v", value)
	}
}

func TestSliceQueue(t *testing.T) {
	queue := NewSliceQueue[string]()
	queue.Enqueue("a")
	queue.Enqueue("b")

	if queue.Size() != 2 {
		t.Errorf("Expected size 2, got %d", queue.Size())
	}

	value, ok := queue.Dequeue()
	if !ok || value != "a" {
		t.Errorf("Expected to dequeue 'a', got %v", value)
	}

	value, ok = queue.Peek()
	if !ok || value != "b" {
		t.Errorf("Expected to peek 'b', got %v", value)
	}
}
