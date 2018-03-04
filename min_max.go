package bst

func (n node) max() int {
	if n.right == nil {
		return n.value
	}
	return n.right.max()
}

func (n node) min() int {
	if n.left == nil {
		return n.value
	}
	return n.left.min()
}

// Max - returns the max value in the tree
func (tree Bst) Max() int {
	return tree.Root.max()
}

// Min - returns the min value in the tree
func (tree Bst) Min() int {
	return tree.Root.min()
}
