package bst

func (n *node) findNode(value int, parent *node) (*node, *node) {
	if n.value == value {
		return n, parent
	}
	if value > n.value {
		if n.right != nil {
			return n.right.findNode(value, n)
		}
		return nil, parent
	}

	if n.left != nil {
		return n.left.findNode(value, n)
	}
	return nil, parent
}

func (n *node) replaceChild(child, newChild *node) bool {
	if n.left == child {
		n.left = newChild
		return true
	} else if n.right == child {
		n.right = newChild
		return true
	}
	return false
}

func (n *node) copyValues(other *node) {
	n.value = other.value
	n.left = other.left
	n.right = other.right
	n.count = other.count
}

func (n *node) deleteSelf(parent *node) bool {
	// Node is a leaf
	if n.right == nil && n.left == nil {
		if parent == nil {
			n.initialized = false
			return true
		}
		return parent.replaceChild(n, nil)
	}
	// Node is a transitional node
	if n.right != nil && n.left != nil {
		// Find the successor (min to the right)
		successor, successorParent := n.right.findMin(n)

		// Copy successor values to the node that is getting deleted
		n.value = successor.value
		n.count = successor.count

		if successorParent != n {
			successorParent.left = successor.right
		} else {
			n.right = nil
		}
		return true
	}
	if n.right != nil {
		if parent == nil {
			n.copyValues(n.right)
			return true
		}
		parent.replaceChild(n, n.right)
		return true
	}
	if n.left != nil {
		if parent == nil {
			n.copyValues(n.left)
			return true
		}
		parent.replaceChild(n, n.left)
	}
	return true
}

func (n *node) delete(value int) bool {
	tbd, tbdParent := n.findNode(value, nil)
	if tbd == nil {
		return false
	}

	tbd.count--

	if tbd.count == 0 {
		return tbd.deleteSelf(tbdParent)
	}
	return true
}

func (n *node) findMin(parent *node) (*node, *node) {
	if n.left == nil {
		return n, parent
	}
	return n.left.findMin(n)
}

// Delete - removes a node from the tree
func (tree *Bst) Delete(value int) error {
	tree.Root.delete(value)
	return nil
}
