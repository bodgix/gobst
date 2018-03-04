package bst

type node struct {
	left        *node
	right       *node
	value       int
	count       int
	initialized bool
}

type Bst struct {
	Root *node
}

// NewBst - return a new BST
func NewBst() *Bst {
	root := &node{}
	return &Bst{root}
}
