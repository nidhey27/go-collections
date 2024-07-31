package collections

import (
	"testing"
)

func TestLinkedList_Add(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	if ll.Size != 3 {
		t.Errorf("Expected size 3, got %d", ll.Size)
	}

	if !ll.Contains(1) || !ll.Contains(2) || !ll.Contains(3) {
		t.Errorf("LinkedList should contain 1, 2, and 3")
	}
}

func TestLinkedList_Remove(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)
	ll.Add(3)

	removed := ll.Remove(2)
	if !removed {
		t.Errorf("Expected to remove 2 from LinkedList")
	}

	if ll.Size != 2 {
		t.Errorf("Expected size 2 after removing 2, got %d", ll.Size)
	}

	if ll.Contains(2) {
		t.Errorf("LinkedList should not contain 2 after removal")
	}
}

func TestLinkedList_RemoveNonExistent(t *testing.T) {
	ll := LinkedList[int]{}
	ll.Add(1)
	ll.Add(2)

	removed := ll.Remove(3)
	if removed {
		t.Errorf("Expected not to remove 3 from LinkedList as it does not exist")
	}

	if ll.Size != 2 {
		t.Errorf("Expected size 2 after attempting to remove non-existent value, got %d", ll.Size)
	}
}

func TestLinkedList_Contains(t *testing.T) {
	ll := LinkedList[string]{}
	ll.Add("a")
	ll.Add("b")
	ll.Add("c")

	if !ll.Contains("a") || !ll.Contains("b") || !ll.Contains("c") {
		t.Errorf("LinkedList should contain 'a', 'b', and 'c'")
	}

	if ll.Contains("d") {
		t.Errorf("LinkedList should not contain 'd'")
	}
}

func TestLinkedList_EmptyList(t *testing.T) {
	ll := LinkedList[int]{}

	if ll.Size != 0 {
		t.Errorf("Expected size 0 for an empty LinkedList, got %d", ll.Size)
	}

	if ll.Contains(1) {
		t.Errorf("LinkedList should not contain any elements")
	}
}
