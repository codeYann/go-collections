package rbtree

import (
	"testing"
)

func TestRBTreeInsertAndSearch(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)

	if tree.Root == tree.Nil || tree.Root.Val != 10 || tree.Root.Color != 'B' {
		t.Errorf("Root node is incorrect")
	}

	tree.Insert(5)
	tree.Insert(30)

	node := tree.Search(5)
	if node == tree.Nil || node.Val != 5 {
		t.Errorf("Expected to find node with value 20, got %v", node)
	}

	if node.Color != 'R' {
		t.Errorf("Expected node with value 20 to be Red, got %c", node.Color)
	}

	node = tree.Search(30)
	if node == tree.Nil || node.Val != 30 {
		t.Errorf("Expected to find node with value 30, got %v", node)
	}

	if node.Color != 'R' {
		t.Errorf("Expected node with value 30 to be Red, got %c", node.Color)
	}

	tree.Insert(45)

	node = tree.Search(45)
	if node == tree.Nil || node.Val != 45 {
		t.Errorf("Expected to find node with value 45, got %v", node)
	}

	if node.Color != 'R' {
		t.Errorf("Expected node with value 45 to be Red, got %c", node.Color)
	}

	if node.Parent.Color != 'B' {
		t.Errorf("Expected parent of node with value 45 to be Black, got %c", node.Parent.Color)
	}

	if tree.Root.Color != 'B' {
		t.Errorf("Expected root node to be Black, got %c", tree.Root.Color)
	}

	tree.Insert(46)

	node = tree.Search(46)
	if node == tree.Nil || node.Val != 46 {
		t.Errorf("Expected to find node with value 46, got %v", node)
	}

	if node.Color != 'R' {
		t.Errorf("Expected node with value 46 to be Red, got %c", node.Color)
	}

	if node.Parent.Color != 'B' {
		t.Errorf("Expected parent of node with value 46 to be Black, got %c", node.Parent.Color)
	}

	if node.Parent.Left.Val != 30 || node.Parent.Left.Color != 'R' {
		t.Errorf("Expected left child of parent of node with value 46 to be Red, got %c", node.Parent.Left.Color)
	}

	if tree.Root.Color != 'B' {
		t.Errorf("Expected root node to be Black, got %c", tree.Root.Color)
	}

	tree.Insert(29)
	node = tree.Search(29)
	if node == tree.Nil || node.Val != 29 {
		t.Errorf("Expected to find node with value 29, got %v", node)
	}

	if node.Color != 'R' {
		t.Errorf("Expected node with value 29 to be Red, got %c", node.Color)
	}

	if node.Parent.Color != 'B' {
		t.Errorf("Expected parent of node with value 29 to be Black, got %c", node.Parent.Color)
	}
}
