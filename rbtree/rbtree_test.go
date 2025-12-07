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

func TestRBTreeRemove(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := NewTree(cmp)

	// Test removing from empty tree
	tree.Remove(10)
	if tree.Root != tree.Nil {
		t.Errorf("Expected empty tree after removing from empty tree")
	}

	// Build a tree with multiple nodes
	values := []int{20, 10, 30, 5, 15, 25, 35}
	for _, v := range values {
		tree.Insert(v)
	}

	// Test removing a leaf node
	tree.Remove(5)
	node := tree.Search(5)
	if node != tree.Nil {
		t.Errorf("Expected node with value 5 to be removed")
	}

	// Verify tree structure is still valid
	if tree.Root.Color != 'B' {
		t.Errorf("Expected root to be black after removing leaf")
	}
	node = tree.Search(10)
	if node == tree.Nil || node.Val != 10 {
		t.Errorf("Expected to find node with value 10 after removing leaf")
	}
	if node.Color != 'B' {
		t.Errorf("Expected node with value 10 to be Black, got %c", node.Color)
	}
	if node.Left != tree.Nil {
		t.Errorf("Expected left child of node with value 10 to be Nil")
	}
	if node.Right == tree.Nil || node.Right.Val != 15 {
		t.Errorf("Expected right child of node with value 10 to be 15")
	}

	// Test removing a node with one child
	tree.Remove(10)
	node = tree.Search(10)
	if node != tree.Nil {
		t.Errorf("Expected node with value 10 to be removed")
	}

	node = tree.Search(15)
	if node == tree.Nil || node.Val != 15 {
		t.Errorf("Expected to find node with value 15 after removing 10")
	}
	if node.Color != 'B' {
		t.Errorf("Expected node with value 15 to be Black, got %c", node.Color)
	}
	if !(node.Left == tree.Nil && node.Right == tree.Nil) {
		t.Errorf("Expected node with value 15 to have no children")
	}

	if tree.Root.Left != node {
		t.Errorf("Expected left child of root to be node with value 15")
	}

	// Test removing a node with two children
	tree.Remove(30)
	node = tree.Search(30)
	if node != tree.Nil {
		t.Errorf("Expected node with value 30 to be removed")
	}

	node = tree.Search(35)
	if node == tree.Nil || node.Val != 35 {
		t.Errorf("Expected to find node with value 35 after removing 30")
	}

	if node.Color != 'B' {
		t.Errorf("Expected node with value 35 to be Black, got %c", node.Color)
	}

	if tree.Root.Right != node {
		t.Errorf("Expected right child of root to be node with value 35")
	}

	if node.Right != tree.Nil {
		t.Errorf("Expected node with value 35 to have no right child")
	}

	if node.Left == tree.Nil || node.Left.Val != 25 {
		t.Errorf("Expected left child of node with value 35 to be 25")
	}

	// Verify remaining nodes are still searchable
	expectedRemaining := []int{20, 15, 35, 25}
	for _, v := range expectedRemaining {
		node := tree.Search(v)
		if node == tree.Nil || node.Val != v {
			t.Errorf("Expected to find node with value %d after removals", v)
		}
	}

	// Test removing root node
	tree.Remove(20)
	node = tree.Search(20)
	if node != tree.Nil {
		t.Errorf("Expected node with value 20 to be removed")
	}

	node = tree.Search(25)

	if node == tree.Nil || node.Val != 25 {
		t.Errorf("Expected to find node with value 25 after removing root")
	}

	if node.Left == tree.Nil || node.Left.Val != 15 {
		t.Errorf("Expected left child of new root to be 15")
	}

	if node.Right == tree.Nil || node.Right.Val != 35 {
		t.Errorf("Expected right child of new root to be 35")
	}

	if node.Right.Color != 'B' {
		t.Errorf("Expected right child of new root to be Black, got %c", node.Right.Color)
	}

	// Verify root is still black
	if tree.Root != tree.Nil && tree.Root.Color != 'B' {
		t.Errorf("Expected root to be black after removing old root")
	}
}

func TestRBTreeRemoveAll(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := NewTree(cmp)

	// Insert multiple values
	values := []int{50, 25, 75, 10, 30, 60, 80}
	for _, v := range values {
		tree.Insert(v)
	}

	// Remove all values
	for _, v := range values {
		tree.Remove(v)
		node := tree.Search(v)
		if node != tree.Nil {
			t.Errorf("Expected node with value %d to be removed", v)
		}
	}

	// Verify tree is empty
	if tree.Root != tree.Nil {
		t.Errorf("Expected tree to be empty after removing all nodes")
	}
}
