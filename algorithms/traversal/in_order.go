package traversal

type Node[T any] struct {
	Val   T
	Left  *Node[T]
	Right *Node[T]
}

type Tree[T any] struct {
	Root *Node[T]
}

func (t *Tree[T]) InOrder(node *Node[T], out *[]T) {
	if node != nil {
		t.InOrder(node.Left, out)
		*out = append(*out, node.Val)
		t.InOrder(node.Right, out)
	}
}
