package collections

import "testing"

func TestLinkedListStack(t *testing.T) {
	stack := NewLinkedListStack[int]()
	stack.Push(1)
	stack.Push(2)

	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}

	value, ok := stack.Pop()
	if !ok || value != 2 {
		t.Errorf("Expected to pop 2, got %v", value)
	}

	value, ok = stack.Peek()
	if !ok || value != 1 {
		t.Errorf("Expected to peek 1, got %v", value)
	}
}

func TestSliceStack(t *testing.T) {
	stack := NewSliceStack[string]()
	stack.Push("x")
	stack.Push("y")

	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}

	value, ok := stack.Pop()
	if !ok || value != "y" {

		value, ok := stack.Pop()
		if !ok || value != "y" {
			t.Errorf("Expected to pop 'y', got %v", value)
		}

		value, ok = stack.Peek()
		if !ok || value != "x" {
			t.Errorf("Expected to peek 'x', got %v", value)
		}

		if stack.Size() != 1 {
			t.Errorf("Expected size 1 after pop, got %d", stack.Size())
		}
	}
}
