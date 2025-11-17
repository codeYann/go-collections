package linkedlist

import (
	"testing"
)

func TestCreateLinkedList(t *testing.T) {
	list := CreateLinkedList[int]()
	if list == nil {
		t.Fatal("CreateLinkedList returned nil")
	}
	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}
	if list.head != nil {
		t.Error("Expected head to be nil")
	}
	if list.tail != nil {
		t.Error("Expected tail to be nil")
	}
}

func TestSize(t *testing.T) {
	list := CreateLinkedList[string]()

	if list.Size() != 0 {
		t.Errorf("Expected size 0, got %d", list.Size())
	}

	list.Insert("first")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}

	list.Append("second")
	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}

	list.Remove("first")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
}

func TestInsert(t *testing.T) {
	list := CreateLinkedList[int]()

	list.Insert(1)
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	if list.head == nil || list.head.Key != 1 {
		t.Error("Head should be 1")
	}
	if list.tail == nil || list.tail.Key != 1 {
		t.Error("Tail should be 1")
	}
	if list.head != list.tail {
		t.Error("Head and tail should be the same for single element")
	}

	list.Insert(2)
	list.Insert(3)

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	if list.head.Key != 3 {
		t.Errorf("Expected head to be 3, got %d", list.head.Key)
	}
	if list.tail.Key != 1 {
		t.Errorf("Expected tail to be 1, got %d", list.tail.Key)
	}

	if list.head.next.Key != 2 {
		t.Errorf("Expected second element to be 2, got %d", list.head.next.Key)
	}
	if list.head.next.next.Key != 1 {
		t.Errorf("Expected third element to be 1, got %d", list.head.next.next.Key)
	}

	if list.head.next.prev != list.head {
		t.Error("Prev pointer not set correctly")
	}
	if list.tail.prev.Key != 2 {
		t.Errorf("Expected tail prev to be 2, got %d", list.tail.prev.Key)
	}
}

func TestAppend(t *testing.T) {
	list := CreateLinkedList[string]()

	list.Append("first")
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	if list.head == nil || list.head.Key != "first" {
		t.Error("Head should be 'first'")
	}
	if list.tail == nil || list.tail.Key != "first" {
		t.Error("Tail should be 'first'")
	}
	if list.head != list.tail {
		t.Error("Head and tail should be the same for single element")
	}

	list.Append("second")
	list.Append("third")

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	if list.head.Key != "first" {
		t.Errorf("Expected head to be 'first', got %s", list.head.Key)
	}
	if list.tail.Key != "third" {
		t.Errorf("Expected tail to be 'third', got %s", list.tail.Key)
	}

	if list.head.next.Key != "second" {
		t.Errorf("Expected second element to be 'second', got %s", list.head.next.Key)
	}
	if list.head.next.next.Key != "third" {
		t.Errorf("Expected third element to be 'third', got %s", list.head.next.next.Key)
	}

	if list.tail.prev.Key != "second" {
		t.Errorf("Expected tail prev to be 'second', got %s", list.tail.prev.Key)
	}
	if list.head.next.prev != list.head {
		t.Error("Prev pointer not set correctly")
	}
}

func TestSearch(t *testing.T) {
	list := CreateLinkedList[int]()

	result := list.Search(1)
	if result != nil {
		t.Error("Expected nil for search in empty list")
	}

	list.Insert(1)
	list.Insert(2)
	list.Insert(3)

	result = list.Search(2)
	if result == nil {
		t.Error("Expected to find element 2")
	}
	if result.Key != 2 {
		t.Errorf("Expected to find key 2, got %d", result.Key)
	}

	result = list.Search(3)
	if result == nil || result.Key != 3 {
		t.Error("Expected to find head element 3")
	}

	result = list.Search(1)
	if result == nil || result.Key != 1 {
		t.Error("Expected to find tail element 1")
	}

	result = list.Search(99)
	if result != nil {
		t.Error("Expected nil for non-existent element")
	}
}

func TestRemove(t *testing.T) {
	t.Run("Remove from empty list", func(t *testing.T) {
		list := CreateLinkedList[int]()
		err := list.Remove(1)
		if err == nil {
			t.Error("Expected error when removing from empty list")
		}
		if err.Error() != "element is not in the list" {
			t.Errorf("Expected error message 'element is not in the list', got %s", err.Error())
		}
	})

	t.Run("Remove single element", func(t *testing.T) {
		list := CreateLinkedList[string]()
		list.Insert("only")

		err := list.Remove("only")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if list.Size() != 0 {
			t.Errorf("Expected size 0, got %d", list.Size())
		}
		if list.head != nil {
			t.Error("Expected head to be nil")
		}
		if list.tail != nil {
			t.Error("Expected tail to be nil")
		}
	})

	t.Run("Remove head", func(t *testing.T) {
		list := CreateLinkedList[int]()
		list.Insert(1)
		list.Insert(2)
		list.Insert(3)
		// List: 3 -> 2 -> 1

		err := list.Remove(3)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if list.Size() != 2 {
			t.Errorf("Expected size 2, got %d", list.Size())
		}
		if list.head.Key != 2 {
			t.Errorf("Expected head to be 2, got %d", list.head.Key)
		}
		if list.head.prev != nil {
			t.Error("Head prev should be nil")
		}
		if list.tail.Key != 1 {
			t.Errorf("Expected tail to be 1, got %d", list.tail.Key)
		}
	})

	t.Run("Remove tail", func(t *testing.T) {
		list := CreateLinkedList[int]()
		list.Append(1)
		list.Append(2)
		list.Append(3)
		// List: 1 -> 2 -> 3

		err := list.Remove(3)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if list.Size() != 2 {
			t.Errorf("Expected size 2, got %d", list.Size())
		}
		if list.head.Key != 1 {
			t.Errorf("Expected head to be 1, got %d", list.head.Key)
		}
		if list.tail.Key != 2 {
			t.Errorf("Expected tail to be 2, got %d", list.tail.Key)
		}
		if list.tail.next != nil {
			t.Error("Tail next should be nil")
		}
	})

	t.Run("Remove middle element", func(t *testing.T) {
		list := CreateLinkedList[int]()
		list.Append(1)
		list.Append(2)
		list.Append(3)
		list.Append(4)
		// List: 1 -> 2 -> 3 -> 4

		err := list.Remove(2)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if list.Size() != 3 {
			t.Errorf("Expected size 3, got %d", list.Size())
		}
		if list.head.Key != 1 {
			t.Errorf("Expected head to be 1, got %d", list.head.Key)
		}
		if list.tail.Key != 4 {
			t.Errorf("Expected tail to be 4, got %d", list.tail.Key)
		}
		// Verify links: 1 -> 3 -> 4
		if list.head.next.Key != 3 {
			t.Errorf("Expected second element to be 3, got %d", list.head.next.Key)
		}
		if list.head.next.prev.Key != 1 {
			t.Errorf("Expected prev of 3 to be 1, got %d", list.head.next.prev.Key)
		}
	})

	t.Run("Remove non-existent element", func(t *testing.T) {
		list := CreateLinkedList[int]()
		list.Insert(1)
		list.Insert(2)

		err := list.Remove(99)
		if err == nil {
			t.Error("Expected error when removing non-existent element")
		}
		if err.Error() != "element is not in the list" {
			t.Errorf("Expected error message 'element is not in the list', got %s", err.Error())
		}
		if list.Size() != 2 {
			t.Errorf("Expected size to remain 2, got %d", list.Size())
		}
	})
}

func TestIntegration(t *testing.T) {
	list := CreateLinkedList[int]()

	list.Insert(1)
	list.Append(2)
	list.Insert(3)
	list.Append(4)
	// List: 3 -> 1 -> 2 -> 4

	if list.Size() != 4 {
		t.Errorf("Expected size 4, got %d", list.Size())
	}

	node := list.Search(1)
	if node == nil || node.Key != 1 {
		t.Error("Failed to find element 1")
	}

	err := list.Remove(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// List: 3 -> 2 -> 4

	if list.Size() != 3 {
		t.Errorf("Expected size 3, got %d", list.Size())
	}
	if list.head.next.Key != 2 {
		t.Errorf("Expected second element to be 2, got %d", list.head.next.Key)
	}

	err = list.Remove(3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// List: 2 -> 4

	if list.head.Key != 2 {
		t.Errorf("Expected head to be 2, got %d", list.head.Key)
	}

	err = list.Remove(4)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// List: 2

	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
	if list.head.Key != 2 || list.tail.Key != 2 {
		t.Error("Head and tail should both be 2")
	}
}

func TestStringType(t *testing.T) {
	list := CreateLinkedList[string]()
	list.Insert("hello")
	list.Append("world")

	if list.Size() != 2 {
		t.Errorf("Expected size 2, got %d", list.Size())
	}

	node := list.Search("hello")
	if node == nil {
		t.Error("Failed to find 'hello'")
	}

	err := list.Remove("world")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if list.Size() != 1 {
		t.Errorf("Expected size 1, got %d", list.Size())
	}
}
