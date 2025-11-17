package linkedlist

import "errors"

type Node[T comparable] struct {
	Key  T
	next *Node[T]
	prev *Node[T]
}

type LinkedList[T comparable] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func createNode[T comparable](key T) *Node[T] {
	return &Node[T]{
		Key:  key,
		next: nil,
		prev: nil,
	}
}

func CreateLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (l *LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) Search(target T) *Node[T] {
	for pt := l.head; pt != nil; pt = pt.next {
		if pt.Key == target {
			return pt
		}
	}
	return nil
}

func (l *LinkedList[T]) Insert(key T) {
	node := createNode(key)

	if l.head == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}

	node.next = l.head
	l.head.prev = node
	l.head = node
	l.size++
}

func (l *LinkedList[T]) Append(key T) {
	node := createNode(key)

	if l.tail == nil {
		l.head = node
		l.tail = node
		l.size++
		return
	}

	node.prev = l.tail
	l.tail.next = node
	l.tail = node
	l.size++
}

func (l *LinkedList[T]) Remove(target T) error {
	node := l.Search(target)
	if node == nil {
		return errors.New("element is not in the list")
	}

	// Single element list
	if l.head == node && l.tail == node {
		l.head = nil
		l.tail = nil
		l.size--
		return nil
	}

	// Remove head
	if node == l.head {
		l.head = node.next
		l.head.prev = nil
		l.size--
		return nil
	}

	// Remove tail
	if node == l.tail {
		l.tail = node.prev
		l.tail.next = nil
		l.size--
		return nil
	}

	// Remove from middle
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
	return nil
}
