package bst

import "testing"

func TestInsert(t *testing.T) {
	tree := NewBst()
	tree.Insert(10)
	if tree.Max() != 10 {
		t.Fatal("Expecting Max to be 10")
	}

	tree.Insert(20)
	if tree.Max() != 20 {
		t.Fatal("Expecting Max to be 20")
	}

	tree.Insert(8)
	if tree.Min() != 8 {
		t.Fatal("Expected Min to be 8")
	}

	tree.Insert(5)
	if tree.Min() != 5 {
		t.Fatal("Expected Min to be 5")
	}
}

func TestDeleteLeaf(t *testing.T) {
	tree := NewBst()
	tree.Insert(10)
	tree.Insert(5)
	tree.Insert(20)

	tree.Delete(20)
	if m := tree.Max(); m != 10 {
		t.Fatal("Expected max to be 10, got ", m)
	}

	tree.Insert(30)
	if m := tree.Max(); m != 30 {
		t.Fatal("Expected max to be 30, got ", m)
	}

	tree.Delete(5)
	if m := tree.Min(); m != 10 {
		t.Fatal("Expected Min to be 10, got ", m)
	}
}

func TestDeleteNodeWithOneChild(t *testing.T) {
	tree := NewBst()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)

	tree.Delete(20)
	if m := tree.Max(); m != 15 {
		t.Fatal("Expected max to be 15, got ", m)
	}
}

func TestDeleteNodeWithTwoChildren(t *testing.T) {
	tree := NewBst()

	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(30)

	tree.Delete(20)
	if m := tree.Max(); m != 30 {
		t.Fatal("Expected max to be 30, got ", m)
	}

	if m := tree.Min(); m != 10 {
		t.Fatal("Expected Min to be 10, got ", m)
	}
}

func TestDeleteRoot(t *testing.T) {
	tree := NewBst()
	tree.Insert(5)
	tree.Insert(9)
	tree.Insert(7)
	tree.Insert(3)
	tree.Insert(4)

	max := tree.Max()
	if max != 9 {
		t.Fatal()
	}

	tree.Delete(5)
	tree.Insert(10)

	max = tree.Max()
	if max != 10 {
		t.Fatal()
	}
}

func TestDeleteTheLastNode(t *testing.T) {
	tree := NewBst()
	tree.Insert(1)
	tree.Delete(1)

	if tree.Root.initialized != false {
		t.Fatal("Expected an empty tree")
	}

	tree.Insert(1)
	tree.Insert(2)
	tree.Delete(1)

	if tree.Root.value != 2 {
		t.Fatal("Expected 2 to be the root but got ", tree.Root.value)
	}

}

func TestDeleteRootWithOneChild(t *testing.T) {
	tree := NewBst()

	tree.Insert(10)
	tree.Insert(20)

	tree.Delete(10)

	max := tree.Max()
	min := tree.Min()

	if max != min || max != 20 {
		t.Fail()
	}

	tree.Insert(5)
	tree.Delete(20)

	max = tree.Max()
	min = tree.Min()

	if max != min || max != 5 {
		t.Fatal()
	}
}

func TestComplexDelete(t *testing.T) {
	tree := NewBst()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(30)
	tree.Insert(12)
	tree.Insert(17)

	tree.Delete(15)

}
