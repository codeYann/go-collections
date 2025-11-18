package stack

import (
	"errors"
)

type Stack[T any] struct {
	capacity uint
	head     int
	arr      []T
}

func NewStack[T any](size uint) *Stack[T] {
	return &Stack[T]{
		capacity: size,
		head:     -1,
		arr:      make([]T, size),
	}
}

func (s Stack[T]) IsEmpty() bool {
	return s.head == -1
}

func (s Stack[T]) IsFull() bool {
	return s.head == int(s.capacity-1)
}

func (s *Stack[T]) Push(elem T) error {
	if s.IsFull() {
		return errors.New("stack is full")
	}

	s.head = s.head + 1
	s.arr[s.head] = elem
	return nil
}

func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	elem := s.arr[s.head]
	s.head = s.head - 1
	return elem, nil
}

func (s Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("stack is empty")
	}

	return s.arr[s.head], nil
}
