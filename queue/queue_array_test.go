package queue

import (
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue, _ := NewQueue[int](5)
	if queue == nil {
		t.Fatal("NewQueue returned nil")
	}
	if !queue.IsEmpty() {
		t.Error("Expected queue to be empty initially")
	}
	if queue.IsFull() {
		t.Error("Expected queue not to be full initially")
	}
	if queue.capacity != 5 {
		t.Errorf("Expected capacity 5, got %d", queue.capacity)
	}
	if queue.front != 0 {
		t.Errorf("Expected front to be 0, got %d", queue.front)
	}
	if queue.rear != 0 {
		t.Errorf("Expected rear to be 0, got %d", queue.rear)
	}
	if len(queue.arr) != 5 {
		t.Errorf("Expected array length 5, got %d", len(queue.arr))
	}

	queue, err := NewQueue[int](0)
	if err == nil {
		t.Fatal("Expected error when creating queue with size 0")
	}
	if queue != nil {
		t.Error("Expected nil queue when size is 0")
	}
}

func TestIsEmpty(t *testing.T) {
	queue, _ := NewQueue[int](3)
	if !queue.IsEmpty() {
		t.Error("Expected empty queue to return true for IsEmpty")
	}

	queue.Enqueue(1)
	if queue.IsEmpty() {
		t.Error("Expected non-empty queue to return false for IsEmpty")
	}

	queue.Dequeue()
	if !queue.IsEmpty() {
		t.Error("Expected empty queue (after dequeue) to return true for IsEmpty")
	}
}

func TestIsFull(t *testing.T) {
	queue, _ := NewQueue[int](3)
	if queue.IsFull() {
		t.Error("Expected non-full queue to return false for IsFull")
	}

	// Circular queue can only hold capacity-1 elements
	queue.Enqueue(1)
	queue.Enqueue(2)
	if !queue.IsFull() {
		t.Error("Expected full queue (capacity-1 elements) to return true for IsFull")
	}

	queue.Dequeue()
	if queue.IsFull() {
		t.Error("Expected non-full queue (after dequeue) to return false for IsFull")
	}
}

func TestEnqueue(t *testing.T) {
	queue, _ := NewQueue[int](3)

	// Enqueue first element
	err := queue.Enqueue(10)
	if err != nil {
		t.Errorf("Unexpected error enqueueing first element: %v", err)
	}
	if queue.IsEmpty() {
		t.Error("Queue should not be empty after enqueue")
	}
	if queue.rear != 1 {
		t.Errorf("Expected rear to be 1 after first enqueue, got %d", queue.rear)
	}
	if queue.front != 0 {
		t.Errorf("Expected front to be 0, got %d", queue.front)
	}

	// Enqueue second element
	err = queue.Enqueue(20)
	if err != nil {
		t.Errorf("Unexpected error enqueueing second element: %v", err)
	}
	if queue.rear != 2 {
		t.Errorf("Expected rear to be 2 after second enqueue, got %d", queue.rear)
	}

	// Try to enqueue when full (capacity-1 = 2 elements)
	err = queue.Enqueue(30)
	if err == nil {
		t.Error("Expected error when enqueueing to full queue")
	}
	if err.Error() != "queue is full" {
		t.Errorf("Expected error message 'queue is full', got %s", err.Error())
	}
}

func TestDequeue(t *testing.T) {
	queue, _ := NewQueue[int](3)

	// Try to dequeue from empty queue
	_, err := queue.Dequeue()
	if err == nil {
		t.Error("Expected error when dequeueing from empty queue")
	}
	if err.Error() != "queue is empty" {
		t.Errorf("Expected error message 'queue is empty', got %s", err.Error())
	}

	// Enqueue and dequeue single element
	queue.Enqueue(10)
	elem, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error dequeueing: %v", err)
	}
	if elem != 10 {
		t.Errorf("Expected dequeued element to be 10, got %d", elem)
	}
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeueing last element")
	}
	if queue.front != 1 {
		t.Errorf("Expected front to be 1 after dequeue, got %d", queue.front)
	}

	// Enqueue multiple elements and dequeue in FIFO order
	queue.Enqueue(1)
	queue.Enqueue(2)

	elem, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 1 {
		t.Errorf("Expected dequeued element to be 1 (FIFO), got %d", elem)
	}
	if queue.front != 2 {
		t.Errorf("Expected front to be 2 after dequeue, got %d", queue.front)
	}

	elem, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 2 {
		t.Errorf("Expected dequeued element to be 2, got %d", elem)
	}
}

func TestPeek(t *testing.T) {
	queue, _ := NewQueue[int](3)

	// Try to peek from empty queue
	_, err := queue.Peek()
	if err == nil {
		t.Error("Expected error when peeking from empty queue")
	}
	if err.Error() != "queue is empty" {
		t.Errorf("Expected error message 'queue is empty', got %s", err.Error())
	}

	// Enqueue and peek
	queue.Enqueue(10)
	elem, err := queue.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking: %v", err)
	}
	if elem != 10 {
		t.Errorf("Expected peeked element to be 10, got %d", elem)
	}
	// Queue should still have the element
	if queue.IsEmpty() {
		t.Error("Queue should not be empty after peek")
	}

	// Enqueue more and peek (should still see first element)
	queue.Enqueue(20)
	elem, err = queue.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking: %v", err)
	}
	if elem != 10 {
		t.Errorf("Expected peeked element to be 10 (FIFO), got %d", elem)
	}
}

func TestFIFOOrder(t *testing.T) {
	queue, _ := NewQueue[int](5)

	// Enqueue multiple elements
	values := []int{1, 2, 3, 4}
	for _, v := range values {
		err := queue.Enqueue(v)
		if err != nil {
			t.Errorf("Unexpected error enqueueing %d: %v", v, err)
		}
	}

	// Dequeue should return in FIFO order (1, 2, 3, 4)
	for i, expected := range values {
		elem, err := queue.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error dequeueing: %v", err)
		}
		if elem != expected {
			t.Errorf("Expected dequeued element at position %d to be %d, got %d", i, expected, elem)
		}
	}

	// Verify empty
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}

func TestCircularWrapping(t *testing.T) {
	queue, _ := NewQueue[int](3)

	// Fill queue (capacity-1 = 2 elements)
	queue.Enqueue(1)
	queue.Enqueue(2)

	// Dequeue one element
	elem, _ := queue.Dequeue()
	if elem != 1 {
		t.Errorf("Expected 1, got %d", elem)
	}

	// Now we have space, enqueue should wrap around
	err := queue.Enqueue(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Verify FIFO order is maintained
	elem, _ = queue.Dequeue()
	if elem != 2 {
		t.Errorf("Expected 2, got %d", elem)
	}
	elem, _ = queue.Dequeue()
	if elem != 3 {
		t.Errorf("Expected 3, got %d", elem)
	}

	// Queue should be empty
	if !queue.IsEmpty() {
		t.Error("Queue should be empty")
	}
}

func TestIntegration(t *testing.T) {
	queue, _ := NewQueue[int](5)

	// Enqueue multiple elements
	values := []int{1, 2, 3, 4}
	for _, v := range values {
		err := queue.Enqueue(v)
		if err != nil {
			t.Errorf("Unexpected error enqueueing %d: %v", v, err)
		}
	}

	// Verify full (capacity-1 = 4 elements)
	if !queue.IsFull() {
		t.Error("Queue should be full")
	}

	// Peek at front
	front, err := queue.Peek()
	if err != nil {
		t.Errorf("Unexpected error peeking: %v", err)
	}
	if front != 1 {
		t.Errorf("Expected front element to be 1, got %d", front)
	}

	// Dequeue all elements in FIFO order
	for i, expected := range values {
		elem, err := queue.Dequeue()
		if err != nil {
			t.Errorf("Unexpected error dequeueing: %v", err)
		}
		if elem != expected {
			t.Errorf("Expected dequeued element at position %d to be %d, got %d", i, expected, elem)
		}
	}

	// Verify empty
	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeuing all elements")
	}
}

func TestEnqueueDequeueSequence(t *testing.T) {
	queue, _ := NewQueue[int](4)

	// Enqueue, dequeue, enqueue, dequeue sequence
	queue.Enqueue(1)
	queue.Enqueue(2)
	elem, _ := queue.Dequeue()
	if elem != 1 {
		t.Errorf("Expected 1, got %d", elem)
	}

	queue.Enqueue(3)
	queue.Enqueue(4)
	elem, _ = queue.Dequeue()
	if elem != 2 {
		t.Errorf("Expected 2, got %d", elem)
	}
	elem, _ = queue.Dequeue()
	if elem != 3 {
		t.Errorf("Expected 3, got %d", elem)
	}
	elem, _ = queue.Dequeue()
	if elem != 4 {
		t.Errorf("Expected 4, got %d", elem)
	}
}

func TestStringType(t *testing.T) {
	queue, _ := NewQueue[string](3)

	queue.Enqueue("hello")
	queue.Enqueue("world")

	if queue.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	elem, err := queue.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != "hello" {
		t.Errorf("Expected 'hello', got %s", elem)
	}

	elem, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != "hello" {
		t.Errorf("Expected 'hello', got %s", elem)
	}

	elem, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != "world" {
		t.Errorf("Expected 'world', got %s", elem)
	}
}

func TestCapacityOne(t *testing.T) {
	queue, _ := NewQueue[int](1)

	if !queue.IsEmpty() {
		t.Error("Queue should be empty initially")
	}

	// With capacity 1, queue can hold 0 elements (capacity-1)
	if !queue.IsFull() {
		t.Error("Queue with capacity 1 should be full immediately (can hold 0 elements)")
	}

	// Try to enqueue
	err := queue.Enqueue(42)
	if err == nil {
		t.Error("Expected error when enqueueing to queue with capacity 1")
	}
	if err.Error() != "queue is full" {
		t.Errorf("Expected error message 'queue is full', got %s", err.Error())
	}
}

func TestCapacityTwo(t *testing.T) {
	queue, _ := NewQueue[int](2)

	if !queue.IsEmpty() {
		t.Error("Queue should be empty initially")
	}

	// With capacity 2, queue can hold 1 element (capacity-1)
	queue.Enqueue(42)
	if !queue.IsFull() {
		t.Error("Queue with capacity 2 should be full after 1 element")
	}

	elem, err := queue.Peek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 42 {
		t.Errorf("Expected 42, got %d", elem)
	}

	elem, err = queue.Dequeue()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if elem != 42 {
		t.Errorf("Expected 42, got %d", elem)
	}

	if !queue.IsEmpty() {
		t.Error("Queue should be empty after dequeue")
	}
}

func TestMultipleWraps(t *testing.T) {
	queue, _ := NewQueue[int](3)

	// Fill and empty multiple times to test wrapping
	for round := 0; round < 3; round++ {
		queue.Enqueue(round*10 + 1)
		queue.Enqueue(round*10 + 2)

		elem1, _ := queue.Dequeue()
		if elem1 != round*10+1 {
			t.Errorf("Round %d: Expected %d, got %d", round, round*10+1, elem1)
		}

		queue.Enqueue(round*10 + 3)

		elem2, _ := queue.Dequeue()
		if elem2 != round*10+2 {
			t.Errorf("Round %d: Expected %d, got %d", round, round*10+2, elem2)
		}

		elem3, _ := queue.Dequeue()
		if elem3 != round*10+3 {
			t.Errorf("Round %d: Expected %d, got %d", round, round*10+3, elem3)
		}

		if !queue.IsEmpty() {
			t.Errorf("Round %d: Queue should be empty", round)
		}
	}
}
