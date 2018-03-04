package bst

import "testing"

func TestPrintPostorder(t *testing.T) {
	tree := NewBst()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(13)
	tree.Insert(17)
	tree.Insert(16)

	PrintPostorder(tree.Root)

	tree.Delete(15)

	PrintPostorder(tree.Root)

	tree.Delete(16)

	PrintPostorder(tree.Root)
}
