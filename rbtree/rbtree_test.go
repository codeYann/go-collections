package rbtree

import (
	"math"
	"math/rand"
	"testing"
)

func generateTestTree() *Tree[int] {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := NewTree(cmp)

	// Manually create nodes for this tree
	/*
	*	 		 	   60
	*	 	    40    70
	*     30  50    100
	 */
	n60 := NewNode(60)
	n40 := NewNode(40)
	n70 := NewNode(70)
	n30 := NewNode(30)
	n50 := NewNode(50)
	n100 := NewNode(100)

	// Setup relationships
	n60.Parent = tree.Nil
	n60.Left = n40
	n60.Right = n70

	n40.Parent = n60
	n40.Left = n30
	n40.Right = n50

	n30.Parent = n40
	n30.Left = tree.Nil
	n30.Right = tree.Nil

	n50.Parent = n40
	n50.Left = tree.Nil
	n50.Right = tree.Nil

	n70.Parent = n60
	n70.Right = n100
	n70.Left = tree.Nil

	n100.Parent = n70
	n100.Left = tree.Nil
	n100.Right = tree.Nil

	tree.Root = n60

	return tree
}

func TestRotateLeft(t *testing.T) {
	// Create a simple tree to test left rotation

	tree := generateTestTree()

	// Perform left rotation on root (60)
	tree.RotateLeft(tree.Root)

	if tree.Root == tree.Nil {
		t.Errorf("Root should not be nil after rotation")
	}

	if tree.Root.Val != 70 {
		t.Errorf("Expected root value to be 70 after left rotation, got %d", tree.Root.Val)
	}

	if tree.Root.Right.Val != 100 {
		t.Errorf("Expected right child of root to be 100 after left rotation, got %d", tree.Root.Right.Val)
	}

	if tree.Root.Left.Val != 60 {
		t.Errorf("Expected left child of root to be 60 after left rotation, got %d", tree.Root.Left.Val)
	}

	if tree.Root.Left.Right != tree.Nil {
		t.Errorf("Expected right child of 60 to be nil after left rotation")
	}

	if tree.Root.Left.Left.Val != 40 {
		t.Errorf("Expected left child of 60 to be 40 after left rotation, got %d", tree.Root.Left.Left.Val)
	}
}

func TestRotateRight(t *testing.T) {
	// Create a simple tree to test right rotation
	tree := generateTestTree()

	// Perform right rotation on root (60)
	tree.RotateRight(tree.Root)

	if tree.Root == tree.Nil {
		t.Errorf("Root should not be nil after rotation")
	}

	if tree.Root.Val != 40 {
		t.Errorf("Expected root value to be 40 after right rotation, got %d", tree.Root.Val)
	}

	if tree.Root.Left.Val != 30 {
		t.Errorf("Expected left child of root to be 30 after right rotation, got %d", tree.Root.Left.Val)
	}

	if tree.Root.Right.Val != 60 {
		t.Errorf("Expected right child of root to be 60 after right rotation, got %d", tree.Root.Right.Val)
	}

	if tree.Root.Right.Left.Val != 50 {
		t.Errorf("Expected left child of 60 to be 50 after right rotation, got %d", tree.Root.Right.Left.Val)
	}

	if tree.Root.Right.Right.Val != 70 {
		t.Errorf("Expected right child of 60 to be 70 after right rotation, got %d", tree.Root.Right.Right.Val)
	}
}

func verifyNode[T any](t *testing.T, tree *Tree[T], node *Node[T], testName string) int {
	if node == tree.Nil {
		return 1
	}

	// Verify if red node has black children
	if node.Color == 'R' {
		if node.Left != tree.Nil && node.Left.Color == 'R' {
			t.Errorf("%s: Red node %v has red left child %v", testName, node.Val, node.Left.Val)
		}

		if node.Right != tree.Nil && node.Right.Color == 'R' {
			t.Errorf("%s: Red node %v has red right child %v", testName, node.Val, node.Right.Val)
		}
	}

	// Verify if all paths from root to leaves have the same number of black nodes
	leftBH := verifyNode(t, tree, node.Left, testName)
	rightBH := verifyNode(t, tree, node.Right, testName)

	if leftBH != rightBH {
		t.Errorf("%s: Black height mismatch at node %v: left BH = %d, right BH = %d", testName, node.Val, leftBH, rightBH)
	}

	if node.Color == 'B' {
		return leftBH + 1
	}

	return leftBH
}

func verifyRBTreeProperties[T any](t *testing.T, tree *Tree[T], testName string) {
	// Verify that the root is black
	if tree.Root != tree.Nil && tree.Root.Color != 'B' {
		t.Errorf("%s: Root must be black", testName)
	}

	// Verify that leaf nodes are black
	if tree.Nil.Color != 'B' {
		t.Errorf("%s: Leaf nodes (NIL) must be black", testName)
	}

	// Verify additional properties recursively
	verifyNode(t, tree, tree.Root, testName)
}

func TestInsertSingleElement(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := NewTree(cmp)

	tree.Insert(10)
	if tree.Root == tree.Nil || tree.Root.Val != 10 {
		t.Errorf("Root should not be nil after insertion and should have value 10")
	}

	// Ensure RB tree properties hold
	verifyRBTreeProperties(t, tree, "TestInsertSingleElement")
}

func TestInsertMultipleElements(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}
	tree := NewTree(cmp)

	elements := []int{10, 20, 30, 40, 50, 60, 70}
	for _, elem := range elements {
		tree.Insert(elem)
	}

	for _, elem := range elements {
		node := tree.Search(elem)
		if node == tree.Nil {
			t.Errorf("Element %d should be found in the tree after insertion", elem)
		}
	}

	// Ensure RB tree properties hold
	verifyRBTreeProperties(t, tree, "TestInsertMultipleElements")
}

func TestInsertFirstCaseRedUncle(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	verifyRBTreeProperties(t, tree, "TestInsertFirstCaseRedUncle")

	node := tree.Search(50)
	if node == tree.Nil || node.Color != 'B' {
		t.Errorf("Node 50 should be black after fixing red uncle case")
	}
	if node.Left == tree.Nil || node.Left.Color != 'B' {
		t.Errorf("Node 30 should be black after fixing red uncle case")
	}
	if node.Right == tree.Nil || node.Right.Color != 'B' {
		t.Errorf("Node 70 should be black after fixing red uncle case")
	}
}

func TestInsertFirstCaseRedUncleSymmetric(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)
	tree.Insert(50)
	tree.Insert(70)
	tree.Insert(30)
	tree.Insert(80)
	verifyRBTreeProperties(t, tree, "TestInsertFirstCaseRedUncleSymmetric")

	node := tree.Search(50)
	if node == tree.Nil || node.Color != 'B' {
		t.Errorf("Node 50 should be black after fixing red uncle case (symmetric)")
	}
	if node.Left == tree.Nil || node.Left.Color != 'B' {
		t.Errorf("Node 30 should be black after fixing red uncle case (symmetric)")
	}
	if node.Right == tree.Nil || node.Right.Color != 'B' {
		t.Errorf("Node 70 should be black after fixing red uncle case (symmetric)")
	}
}

func TestInsertSecondCaseUncleBlack(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(40)
	verifyRBTreeProperties(t, tree, "TestInsertSecondCaseUncleBlack")

	node := tree.Search(40)

	if node == tree.Nil || node.Color != 'B' || node != tree.Root {
		t.Errorf("Node 40 should be black and root after fixing second case with black uncle")
	}

	if node.Left == tree.Nil || node.Left.Val != 30 || node.Left.Color != 'R' {
		t.Errorf("Node 30 should be red and left child of 40 after fixing second case with black uncle")
	}

	if node.Right == tree.Nil || node.Right.Val != 50 || node.Right.Color != 'R' {
		t.Errorf("Node 50 should be red and right child of 40 after fixing second case with black uncle")
	}
}

func TestInsertSecondCaseUncleBlackSymmetric(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)
	tree.Insert(50)
	tree.Insert(70)
	tree.Insert(60)
	verifyRBTreeProperties(t, tree, "TestInsertSecondCaseUncleBlackSymmetric")

	node := tree.Search(60)
	if node == tree.Nil || node.Color != 'B' || node != tree.Root {
		t.Errorf("Node 60 should be black and root after fixing second case with black uncle (symmetric)")
	}

	if node.Left == tree.Nil || node.Left.Val != 50 || node.Left.Color != 'R' {
		t.Errorf("Node 50 should be red and left child of 60 after fixing second case with black uncle (symmetric)")
	}

	if node.Right == tree.Nil || node.Right.Val != 70 || node.Right.Color != 'R' {
		t.Errorf("Node 70 should be red and right child of 60 after fixing second case with black uncle (symmetric)")
	}
}

func TestInsertThirdCaseUncleBlack(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)

	tree.Insert(50)
	tree.Insert(30)
	tree.Insert(20)
	verifyRBTreeProperties(t, tree, "TestInsertThirdCaseUncleBlack")

	node := tree.Search(30)
	if node == tree.Nil || node.Color != 'B' || node != tree.Root {
		t.Errorf("Node 30 should be black and root after fixing third case with black uncle")
	}

	if node.Left == tree.Nil || node.Left.Val != 20 || node.Left.Color != 'R' {
		t.Errorf("Node 20 should be red and left child of 30 after fixing third case with black uncle")
	}

	if node.Right == tree.Nil || node.Right.Val != 50 || node.Right.Color != 'R' {
		t.Errorf("Node 50 should be red and right child of 30 after fixing third case with black uncle")
	}
}

func TestInsertThirdCaseUncleBlackSymmetric(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)
	tree.Insert(50)
	tree.Insert(70)
	tree.Insert(80)
	verifyRBTreeProperties(t, tree, "TestInsertThirdCaseUncleBlackSymmetric")

	node := tree.Search(70)
	if node == tree.Nil || node.Color != 'B' || node != tree.Root {
		t.Errorf("Node 70 should be black and root after fixing third case with black uncle (symmetric)")
	}

	if node.Left == tree.Nil || node.Left.Val != 50 || node.Left.Color != 'R' {
		t.Errorf("Node 50 should be red and left child of 70 after fixing third case with black uncle (symmetric)")
	}

	if node.Right == tree.Nil || node.Right.Val != 80 || node.Right.Color != 'R' {
		t.Errorf("Node 80 should be red and right child of 70 after fixing third case with black uncle (symmetric)")
	}
}

func TestRandomInserts(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)

	for range 100 {
		value := rand.Intn(1000)
		tree.Insert(value)
		verifyRBTreeProperties(t, tree, "TestRandomInserts")
	}
}

func TestFinalTreeHeight(t *testing.T) {
	cmp := func(a, b int) int {
		return a - b
	}

	tree := NewTree(cmp)
	numInserts := 1000
	for i := 0; i < numInserts; i++ {
		value := rand.Intn(10000)
		tree.Insert(value)
	}

	// The height of a red-black tree is at most 2*log2(n+1)
	expectedMaxHeight := 2 * int(math.Log2(float64(numInserts+1)))
	actualHeight := tree.Height(tree.Root)

	if actualHeight > expectedMaxHeight {
		t.Errorf("Tree height %d exceeds expected maximum height %d", actualHeight, expectedMaxHeight)
	}
}
