package collections

import "reflect"

// Node represents a node in a LinkedList.
type Node[T any] struct {
	Value T
	Next  *Node[T]
}

// LinkedList is a generic linked list.
type LinkedList[T any] struct {
	Head *Node[T]
	Size int
}

// Add appends a new value to the end of the LinkedList.
func (ll *LinkedList[T]) Add(value T) {
	newNode := &Node[T]{Value: value}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

// Remove deletes the first occurrence of the value from the LinkedList.
func (ll *LinkedList[T]) Remove(value T) bool {
	if ll.Head == nil {
		return false
	}
	if reflect.DeepEqual(ll.Head.Value, value) {
		ll.Head = ll.Head.Next
		ll.Size--
		return true
	}
	current := ll.Head
	for current.Next != nil && !reflect.DeepEqual(current.Next.Value, value) {
		current = current.Next
	}
	if current.Next == nil {
		return false
	}
	current.Next = current.Next.Next
	ll.Size--
	return true
}

// Contains checks if the value exists in the LinkedList.
func (ll *LinkedList[T]) Contains(value T) bool {
	current := ll.Head
	for current != nil {
		if reflect.DeepEqual(current.Value, value) {
			return true
		}
		current = current.Next
	}
	return false
}
