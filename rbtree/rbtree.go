package rbtree

// Node represents a node in the red-black tree.
// It holds a value of generic type T and pointers to its parent, left, and right children.
// The Color field indicates whether the node is red ('R') or black ('B').
type Node[T any] struct {
	Val    T
	Left   *Node[T]
	Right  *Node[T]
	Parent *Node[T]
	Color  byte // 'R' for Red, 'B' for Black
}

// Comparator defines a function type for comparing two values of type T.
// It returns a negative value if a < b, zero if a == b, and a positive value if a > b.
type Comparator[T any] func(a, b T) int

// Tree represents a red-black tree data structure.
// It maintains the root node, a comparator function, and a sentinel Nil node.
type Tree[T any] struct {
	Root       *Node[T]
	Comparator Comparator[T]
	Nil        *Node[T]
}

// NewNode creates and returns a new red node with the given value.
func NewNode[T any](val T) *Node[T] {
	return &Node[T]{Val: val, Color: 'R'}
}

// NewTree creates and returns a new red-black tree with the specified comparator function.
func NewTree[T any](cmp Comparator[T]) *Tree[T] {
	var Nil *Node[T] = &Node[T]{Color: 'B'} // define a sentinel Nil node
	return &Tree[T]{Root: Nil, Comparator: cmp, Nil: Nil}
}

// RotateLeft performs a left rotation on the given node x.
// This operation maintains the red-black tree properties.
func (t *Tree[T]) RotateLeft(x *Node[T]) {
	y := x.Right
	x.Right = y.Left
	if y.Left != t.Nil {
		y.Left.Parent = x
	}
	y.Parent = x.Parent
	if x.Parent == t.Nil {
		t.Root = y
	} else if x == x.Parent.Left {
		x.Parent.Left = y
	} else {
		x.Parent.Right = y
	}
	y.Left = x
	x.Parent = y
}

// RotateRight performs a right rotation on the given node y.
// This operation maintains the red-black tree properties.
func (t *Tree[T]) RotateRight(y *Node[T]) {
	x := y.Left
	y.Left = x.Right
	if x.Right != t.Nil {
		x.Right.Parent = y
	}

	x.Parent = y.Parent
	if y.Parent == t.Nil {
		t.Root = x
	} else if y == y.Parent.Right {
		y.Parent.Right = x
	} else {
		y.Parent.Left = x
	}

	x.Right = y
	y.Parent = x
}

func (t *Tree[T]) insertFixup(z *Node[T]) {
	for z.Parent != t.Nil && z.Parent.Color == 'R' {
		if z.Parent.Parent == t.Nil {
			break
		}

		if z.Parent == z.Parent.Parent.Left {
			y := z.Parent.Parent.Right
			if y != t.Nil && y.Color == 'R' {
				z.Parent.Color = 'B'
				y.Color = 'B'
				z.Parent.Parent.Color = 'R'
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Right {
					z = z.Parent
					t.RotateLeft(z)
				}
				z.Parent.Color = 'B'
				z.Parent.Parent.Color = 'R'
				t.RotateRight(z.Parent.Parent)
			}
		} else {
			y := z.Parent.Parent.Left
			if y != t.Nil && y.Color == 'R' {
				z.Parent.Color = 'B'
				y.Color = 'B'
				z.Parent.Parent.Color = 'R'
				z = z.Parent.Parent
			} else {
				if z == z.Parent.Left {
					z = z.Parent
					t.RotateRight(z)
				}
				z.Parent.Color = 'B'
				z.Parent.Parent.Color = 'R'
				t.RotateLeft(z.Parent.Parent)
			}
		}
	}
	t.Root.Color = 'B'
}

// Insert adds a new element to the red-black tree while maintaining its properties.
func (t *Tree[T]) Insert(elem T) {
	node := NewNode(elem)

	if t.Root == t.Nil {
		t.Root = node
		t.Root.Color = 'B'
		t.Root.Parent = t.Nil
		t.Root.Left = t.Nil
		t.Root.Right = t.Nil
		return
	}

	current := t.Root

	for {
		cmp := t.Comparator(elem, current.Val)

		if cmp < 0 {
			if current.Left == t.Nil {
				current.Left = node
				node.Parent = current
				node.Left = t.Nil
				node.Right = t.Nil
				break
			}
			current = current.Left
		} else {
			if current.Right == t.Nil {
				current.Right = node
				node.Parent = current
				node.Left = t.Nil
				node.Right = t.Nil
				break
			}
			current = current.Right
		}
	}

	t.insertFixup(node)
}

// Search looks for a node with the specified value in the red-black tree.
// It returns the node if found, or the sentinel Nil node if not found.
func (t *Tree[T]) Search(elem T) *Node[T] {
	current := t.Root

	for current != t.Nil {
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

	return t.Nil
}

// Minimum finds and returns the node with the smallest value in the subtree rooted at the given node.
func (t *Tree[T]) Minimum(node *Node[T]) *Node[T] {
	current := node
	for current.Left != t.Nil {
		current = current.Left
	}
	return current
}

// Maximum finds and returns the node with the largest value in the subtree rooted at the given node.
func (t *Tree[T]) Maximum(node *Node[T]) *Node[T] {
	current := node
	for current.Right != t.Nil {
		current = current.Right
	}
	return current
}

// Height calculates and returns the height of the subtree rooted at the given node.
// The height of a tree with only the Nil node is -1.
func (t *Tree[T]) Height(node *Node[T]) int {
	if node == t.Nil {
		return -1
	}
	return 1 + max(t.Height(node.Left), t.Height(node.Right))
}

// Size calculates and returns the number of nodes in the subtree rooted at the given node.
func (t *Tree[T]) Size(node *Node[T]) int {
	if node == t.Nil {
		return 0
	}
	return 1 + t.Size(node.Left) + t.Size(node.Right)
}

// Successor finds and returns the successor of the given node in the red-black tree.
// The successor is the node with the smallest value greater than the given node's value.
func (t *Tree[T]) Successor(node *Node[T]) *Node[T] {
	if node.Right != t.Nil {
		return t.Minimum(node.Right)
	}

	current := node
	parent := current.Parent
	for parent != t.Nil && current == parent.Right {
		current = parent
		parent = parent.Parent
	}
	return parent
}

// Predecessor finds and returns the predecessor of the given node in the red-black tree.
// The predecessor is the node with the largest value less than the given node's value.
func (t *Tree[T]) Predecessor(node *Node[T]) *Node[T] {
	if node.Left != t.Nil {
		return t.Maximum(node.Left)
	}

	current := node
	parent := current.Parent
	for parent != t.Nil && current == parent.Left {
		current = parent
		parent = parent.Parent
	}
	return parent
}

// Remove deletes a node with the specified value from the red-black tree.
// This method is not yet implemented.
func (t *Tree[T]) Remove(elem T) {
	panic("Not implemented")
}
