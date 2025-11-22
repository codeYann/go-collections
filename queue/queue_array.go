package queue

import "errors"

type Queue[T any] struct {
	capacity uint
	front    int
	rear     int
	arr      []T
}

func NewQueue[T any](size uint) *Queue[T] {
	return &Queue[T]{
		capacity: size,
		front:    0,
		rear:     0,
		arr:      make([]T, size),
	}
}

func (q Queue[T]) IsEmpty() bool {
	return q.front == q.rear
}

func (q Queue[T]) IsFull() bool {
	return (q.rear+1)%int(q.capacity) == q.front
}

func (q *Queue[T]) Enqueue(elem T) error {
	if q.IsFull() {
		return errors.New("queue is full")
	}

	q.arr[q.rear] = elem
	q.rear = (q.rear + 1) % int(q.capacity)
	return nil
}

func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}

	elem := q.arr[q.front]
	q.front = (q.front + 1) % int(q.capacity)
	return elem, nil
}

func (q *Queue[T]) Peek() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("queue is empty")
	}
	return q.arr[q.front], nil
}
