package bst

func (n *node) insert(value int) error {
	if !n.initialized {
		n.count = 1
		n.value = value
		n.initialized = true
		return nil
	}
	if n.value == value {
		n.count++
		return nil
	}
	if value > n.value {
		if n.right == nil {
			n.right = &node{nil, nil, value, 1, true}
			return nil
		}
		return n.right.insert(value)
	}
	if n.left == nil {
		n.left = &node{nil, nil, value, 1, true}
		return nil
	}
	return n.left.insert(value)
}

// Insert - insert a new node into the tree
func (tree *Bst) Insert(value int) error {
	return tree.Root.insert(value)
}
