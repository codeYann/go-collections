package bst

// Node represents a node in a binary search tree.
type Node[T any] struct {
	Val    T
	Left   *Node[T]
	Right  *Node[T]
	Parent *Node[T]
}

// Comparator is a function type that compares two values of type T.
// It returns a negative number if a < b, zero if a == b, and a positive number if a > b.
type Comparator[T any] func(a, b T) int

// Tree represents a binary search tree.
type Tree[T any] struct {
	Root       *Node[T]
	Comparator Comparator[T]
}

// NewNode creates and returns a new node with the given value.
func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val}
}

// NewTree creates and returns a new empty binary search tree.
func NewTree[T any]() *Tree[T] {
	return &Tree[T]{Root: nil}
}

// Insert adds a new element to the binary search tree.
// The element is inserted according to the tree's comparator function.
func (t *Tree[T]) Insert(elem T) {
	newNode := NewNode(elem)

	if t.Root == nil {
		t.Root = newNode
		return
	}

	current := t.Root

	for {
		cmp := t.Comparator(elem, current.Val)

		if cmp < 0 {
			if current.Left == nil {
				current.Left = newNode
				newNode.Parent = current
				return
			}
			current = current.Left
		} else {
			if current.Right == nil {
				current.Right = newNode
				newNode.Parent = current
				return
			}
			current = current.Right
		}
	}
}

// Search searches for an element in the binary search tree.
// Returns the node containing the element, or nil if not found.
func (t *Tree[T]) Search(elem T) *Node[T] {
	current := t.Root

	for current != nil {
		cmp := t.Comparator(elem, current.Val)

		if cmp == 0 {
			return current
		}

		if cmp < 0 {
			current = current.Left
		} else {
			current = current.Right
		}
	}

	return nil
}

// Minimum returns the node with the minimum value in the subtree rooted at the given node.
func (t *Tree[T]) Minimum(node *Node[T]) *Node[T] {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// Maximum returns the node with the maximum value in the subtree rooted at the given node.
func (t *Tree[T]) Maximum(node *Node[T]) *Node[T] {
	current := node
	for current.Right != nil {
		current = current.Right
	}
	return current
}

// Height returns the height of the subtree rooted at the given node.
// The height of a leaf node is 0, and the height of an empty tree is -1.
func (t *Tree[T]) Height(node *Node[T]) int {
	if node == nil {
		return -1
	}

	return max(t.Height(node.Left), t.Height(node.Right)) + 1
}

// Size returns the number of nodes in the subtree rooted at the given node.
func (t *Tree[T]) Size(node *Node[T]) int {
	if node == nil {
		return 0
	}
	return t.Size(node.Left) + t.Size(node.Right) + 1
}

// Successor returns the node with the smallest value greater than the given node's value.
// Returns nil if the node is the maximum node in the tree.
func (t *Tree[T]) Successor(node *Node[T]) *Node[T] {
	if node.Right != nil {
		return t.Minimum(node.Right)
	}

	current := node
	parent := node.Parent

	for parent != nil && current == parent.Right {
		current = parent
		parent = parent.Parent
	}
	return parent
}

// Predecessor returns the node with the largest value smaller than the given node's value.
// Returns nil if the node is the minimum node in the tree.
func (t *Tree[T]) Predecessor(node *Node[T]) *Node[T] {
	if node.Left != nil {
		return t.Maximum(node.Left)
	}

	current := node
	parent := node.Parent
	for parent != nil && current == parent.Left {
		current = parent
		parent = parent.Parent
	}
	return parent
}

// transplant replaces the subtree rooted at node u with the subtree rooted at node v.
func (t *Tree[T]) transplant(u, v *Node[T]) {
	if u.Parent == nil {
		t.Root = v
	} else if u == u.Parent.Left {
		u.Parent.Left = v
	} else {
		u.Parent.Right = v
	}

	if v != nil {
		v.Parent = u.Parent
	}
}

// Remove deletes the node containing the given element from the tree.
// If the element is not found, the tree remains unchanged.
func (t *Tree[T]) Remove(elem T) {
	node := t.Search(elem)
	if node == nil {
		return
	}

	if node.Left == nil {
		t.transplant(node, node.Right)
		return
	}

	if node.Right == nil {
		t.transplant(node, node.Left)
		return
	}

	successor := t.Successor(node)

	if successor.Parent != node {
		t.transplant(successor, successor.Right)
		successor.Right = node.Right
		successor.Right.Parent = successor
	}

	t.transplant(node, successor)
	successor.Left = node.Left
	successor.Left.Parent = successor
}
