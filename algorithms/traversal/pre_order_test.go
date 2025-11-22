package traversal

import (
	"testing"
)

/*
			1
	 2     3

4   5  6    7
*/
func TestPreOrderTraversal(t *testing.T) {
	tree := &Tree[int]{
		Root: &Node[int]{
			Val: 1,
			Left: &Node[int]{
				Val:   2,
				Left:  &Node[int]{Val: 4},
				Right: &Node[int]{Val: 5},
			},
			Right: &Node[int]{
				Val:   3,
				Left:  &Node[int]{Val: 6},
				Right: &Node[int]{Val: 7},
			},
		},
	}

	var result []int
	tree.PreOrder(tree.Root, &result)

	expected := []int{1, 2, 4, 5, 3, 6, 7}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %d at index %d, got %d", v, i, result[i])
		}
	}
}

/*
			a
	 b     c

d   e   f   g
*/
func TestPreOrderTraversalString(t *testing.T) {
	tree := &Tree[string]{
		Root: &Node[string]{
			Val: "a",
			Left: &Node[string]{
				Val:   "b",
				Left:  &Node[string]{Val: "d"},
				Right: &Node[string]{Val: "e"},
			},
			Right: &Node[string]{
				Val:   "c",
				Left:  &Node[string]{Val: "f"},
				Right: &Node[string]{Val: "g"},
			},
		},
	}

	var result []string
	tree.PreOrder(tree.Root, &result)

	expected := []string{"a", "b", "d", "e", "c", "f", "g"}
	for i, v := range expected {
		if result[i] != v {
			t.Errorf("Expected %s at index %d, got %s", v, i, result[i])
		}
	}
}
