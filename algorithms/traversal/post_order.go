package traversal

func (t *Tree[T]) PostOrder(node *Node[T], out *[]T) {
	if node != nil {
		t.PostOrder(node.Left, out)
		t.PostOrder(node.Right, out)
		*out = append(*out, node.Val)
	}
}
