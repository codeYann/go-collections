package traversal

func (t *Tree[T]) PreOrder(node *Node[T], out *[]T) {
	if node != nil {
		*out = append(*out, node.Val)
		t.PreOrder(node.Left, out)
		t.PreOrder(node.Right, out)
	}
}
