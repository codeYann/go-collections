package bst

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	node := tree.Search(7)
	if node == nil || node.Val != 7 {
		t.Errorf("Expected to find node with value 7, got %v", node)
	}

	node = tree.Search(20)
	if node != nil {
		t.Errorf("Expected to not find node with value 20, but got %v", node)
	}
}

func TestMinimumAndMaximum(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	min := tree.Minimum(tree.Root)
	if min == nil || min.Val != 3 {
		t.Errorf("Expected minimum value 3, got %v", min)
	}

	max := tree.Maximum(tree.Root)
	if max == nil || max.Val != 15 {
		t.Errorf("Expected maximum value 15, got %v", max)
	}
}

func TestHeight(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	height := tree.Height(tree.Root)
	if height != 2 {
		t.Errorf("Expected height 2, got %d", height)
	}
}

func TestSize(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	size := tree.Size(tree.Root)
	if size != 5 {
		t.Errorf("Expected size 5, got %d", size)
	}
}

func TestSucessor(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	node := tree.Search(5)
	succ := tree.Sucessor(node)
	if succ == nil || succ.Val != 7 {
		t.Errorf("Expected successor of 5 to be 7, got %v", succ)
	}

	node = tree.Search(10)
	succ = tree.Sucessor(node)
	if succ == nil || succ.Val != 15 {
		t.Errorf("Expected successor of 10 to be 15, got %v", succ)
	}

	node = tree.Search(3)
	succ = tree.Sucessor(node)
	if succ == nil || succ.Val != 5 {
		t.Errorf("Expected successor of 3 to be 5, got %v", succ)
	}

	node = tree.Search(7)
	succ = tree.Sucessor(node)
	if succ == nil || succ.Val != 10 {
		t.Errorf("Expected successor of 7 to be 10, got %v", succ)
	}
}

func TestPredecessor(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	node := tree.Search(7)
	pred := tree.Predecessor(node)
	if pred == nil || pred.Val != 5 {
		t.Errorf("Expected predecessor of 7 to be 5, got %v", pred)
	}

	node = tree.Search(10)
	pred = tree.Predecessor(node)
	if pred == nil || pred.Val != 7 {
		t.Errorf("Expected predecessor of 10 to be 7, got %v", pred)
	}

	node = tree.Search(15)
	pred = tree.Predecessor(node)
	if pred == nil || pred.Val != 10 {
		t.Errorf("Expected predecessor of 15 to be 10, got %v", pred)
	}

	node = tree.Search(3)
	pred = tree.Predecessor(node)
	if pred != nil {
		t.Errorf("Expected predecessor of 3 to be nil, got %v", pred)
	}
}

func TestRemove(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := &Tree[int]{Comparator: cmp}
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(15)
	tree.Insert(3)
	tree.Insert(7)

	tree.Remove(5)
	node := tree.Search(5)
	if node != nil {
		t.Errorf("Expected node with value 5 to be removed, but got %v", node)
	}

	node = tree.Search(7)
	if node == nil || node.Val != 7 {
		t.Errorf("Expected to find node with value 7, got %v", node)
	}

	if node.Left.Val != 3 {
		t.Errorf("Expected left child of node 7 to be 3, got %v", node.Left)
	}

	if node.Left.Parent.Val != 7 {
		t.Errorf("Expected parent of node 3 to be 7, got %v", node.Left.Parent)
	}

	if node.Right != nil {
		t.Errorf("Expected right child of node 7 to be nil, got %v", node.Right)
	}

	if tree.Root.Left.Val != 7 {
		t.Errorf("Expected left child of root to be 7, got %v", tree.Root.Left)
	}

	if tree.Root.Left.Parent.Val != 10 {
		t.Errorf("Expected parent of node 7 to be 10, got %v", tree.Root.Left.Parent)
	}
}
