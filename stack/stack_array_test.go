package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := NewStack[int](5)
	if stack == nil {
		t.Fatal("NewStack returned nil")
	}
	if !stack.IsEmpty() {
		t.Error("Expected stack to be empty initially")
	}
	if stack.IsFull() {
		t.Error("Expected stack not to be full initially")
	}
	if stack.capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", stack.capacity)
	}
	if stack.head != -1 {
		t.Errorf("Expected head to be -1, got %d", stack.head)
	}
	if len(stack.arr) != 5 {
		t.Errorf("Expected array length 5, got %d", len(stack.arr))
	}
}

func TestIsEmpty(t *testing.T) {
	stack := NewStack[int](3)
	if !stack.IsEmpty() {
		t.Error("Expected empty stack to return true for IsEmpty")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Error("Expected non-empty stack to return false for IsEmpty")
	}

	stack.Pop()
	if !stack.IsEmpty() {
		t.Error("Expected empty stack (after pop) to return true for IsEmpty")
	}
}

func TestIsFull(t *testing.T) {
	stack := NewStack[int](3)
	if stack.IsFull() {
		t.Error("Expected non-full stack to return false for IsFull")
	}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	if !stack.IsFull() {
		t.Error("Expected full stack to return true for IsFull")
	}

	stack.Pop()
	if stack.IsFull() {
		t.Error("Expected non-full stack (after pop) to return false for IsFull")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack[int](3)

	// Push first element
	err := stack.Push(10)
	if err != nil {
		t.Errorf("Unexpected error pushing first element: %v", err)
	}
	if stack.IsEmpty() {
		t.Error("Stack should not be empty after push")
	}
	if stack.head != 0 {
		t.Errorf("Expected head to be 0 after first push, got %d", stack.head)
	}

	// Push second element
	err = stack.Push(20)
	if err != nil {
		t.Errorf("Unexpected error pushing second element: %v", err)
	}
	if stack.head != 1 {
		t.Errorf("Expected head to be 1 after second push, got %d", stack.head)
	}

	// Push third element
	err = stack.Push(30)
	if err != nil {
		t.Errorf("Unexpected error pushing third element: %v", err)
	}
	if stack.head != 2 {
		t.Errorf("Expected head to be 2 after third push, got %d", stack.head)
	}
	if !stack.IsFull() {
		t.Error("Stack should be full after pushing capacity elements")
	}

	// Try to push when full
	err = stack.Push(40)
	if err == nil {
		t.Error("Expected error when pushing to full stack")
	}
	if err.Error() != "stack is full" {
		t.Errorf("Expected error message 'stack is full', got %s", err.Error())
	}
}

func TestPop(t *testing.T) {
	stack := NewStack[int](3)

	// Try to pop from empty stack
	_, err := stack.Pop()
	if err == nil {
		t.Error("Expected error when popping from empty stack")
	}
	if err.Error() != "stack is empty" {
		t.Errorf("Expected error message 'stack is empty', got %s", err.Error())
	}

	// Push and pop single element
	stack.Push(10)
	elem, err := stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error popping: %v", err)
	}
	if elem != 10 {
		t.Errorf("Expected popped element to be 10, got %d", elem)
	}
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping last element")
	}
	if stack.head != -1 {
		t.Errorf("Expected head to be -1 after popping last element, got %d", stack.head)
	}

	// Push multiple elements and pop in LIFO order
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	elem, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 3 {
		t.Errorf("Expected popped element to be 3 (LIFO), got %d", elem)
	}
	if stack.head != 1 {
		t.Errorf("Expected head to be 1 after pop, got %d", stack.head)
	}

	elem, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 2 {
		t.Errorf("Expected popped element to be 2, got %d", elem)
	}

	elem, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 1 {
		t.Errorf("Expected popped element to be 1, got %d", elem)
	}
}

func TestPeek(t *testing.T) {
	stack := NewStack[int](3)

	// Try to peek from empty stack
	_, err := stack.Peek()
	if err == nil {
		t.Error("Expected error when peeking from empty stack")
	}
	if err.Error() != "stack is empty" {
		t.Errorf("Expected error message 'stack is empty', got %s", err.Error())
	}

	// Push and peek
	stack.Push(10)
	elem, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking: %v", err)
	}
	if elem != 10 {
		t.Errorf("Expected peeked element to be 10, got %d", elem)
	}
	// Stack should still have the element
	if stack.IsEmpty() {
		t.Error("Stack should not be empty after peek")
	}

	// Push more and peek
	stack.Push(20)
	elem, err = stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking: %v", err)
	}
	if elem != 20 {
		t.Errorf("Expected peeked element to be 20, got %d", elem)
	}
	// Verify stack still has both elements
	if stack.head != 1 {
		t.Errorf("Expected head to be 1, got %d", stack.head)
	}
}

func TestIntegration(t *testing.T) {
	stack := NewStack[int](5)

	// Push multiple elements
	values := []int{1, 2, 3, 4, 5}
	for _, v := range values {
		err := stack.Push(v)
		if err != nil {
			t.Errorf("Unexpected error pushing %d: %v", v, err)
		}
	}

	// Verify full
	if !stack.IsFull() {
		t.Error("Stack should be full")
	}

	// Peek at top
	top, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking: %v", err)
	}
	if top != 5 {
		t.Errorf("Expected top element to be 5, got %d", top)
	}

	// Pop all elements in reverse order
	for i := len(values) - 1; i >= 0; i-- {
		elem, err := stack.Pop()
		if err != nil {
			t.Errorf("Unexpected error popping: %v", err)
		}
		if elem != values[i] {
			t.Errorf("Expected popped element to be %d, got %d", values[i], elem)
		}
	}

	// Verify empty
	if !stack.IsEmpty() {
		t.Error("Stack should be empty after popping all elements")
	}
}

func TestPushPopSequence(t *testing.T) {
	stack := NewStack[int](3)

	// Push, pop, push, pop sequence
	stack.Push(1)
	stack.Push(2)
	elem, _ := stack.Pop()
	if elem != 2 {
		t.Errorf("Expected 2, got %d", elem)
	}

	stack.Push(3)
	stack.Push(4)
	elem, _ = stack.Pop()
	if elem != 4 {
		t.Errorf("Expected 4, got %d", elem)
	}
	elem, _ = stack.Pop()
	if elem != 3 {
		t.Errorf("Expected 3, got %d", elem)
	}
	elem, _ = stack.Pop()
	if elem != 1 {
		t.Errorf("Expected 1, got %d", elem)
	}
}

func TestStringType(t *testing.T) {
	stack := NewStack[string](3)

	stack.Push("hello")
	stack.Push("world")

	if stack.IsEmpty() {
		t.Error("Stack should not be empty")
	}

	elem, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != "world" {
		t.Errorf("Expected 'world', got %s", elem)
	}

	elem, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != "world" {
		t.Errorf("Expected 'world', got %s", elem)
	}

	elem, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != "hello" {
		t.Errorf("Expected 'hello', got %s", elem)
	}
}

func TestSingleElementStack(t *testing.T) {
	stack := NewStack[int](1)

	if !stack.IsEmpty() {
		t.Error("Stack should be empty initially")
	}

	err := stack.Push(42)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !stack.IsFull() {
		t.Error("Stack should be full with capacity 1")
	}

	elem, err := stack.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 42 {
		t.Errorf("Expected 42, got %d", elem)
	}

	elem, err = stack.Pop()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 42 {
		t.Errorf("Expected 42, got %d", elem)
	}

	if !stack.IsEmpty() {
		t.Error("Stack should be empty after pop")
	}
}
