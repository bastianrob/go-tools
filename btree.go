package gotools

type Comparable interface {
	LessThan(Comparable) bool
	Equals(Comparable) bool
}

type BTree struct {
	Value Comparable
	Left  *BTree
	Right *BTree
}

// Insert adds a new node to the binary tree
func (n *BTree) Insert(value Comparable) *BTree {
	if n == nil {
		return &BTree{Value: value}
	}

	if value.LessThan(n.Value) {
		n.Left = n.Left.Insert(value)
	} else {
		n.Right = n.Right.Insert(value)
	}

	return n
}

// Search searches the binary search tree for a value
func (n *BTree) Search(value Comparable) *BTree {
	if n == nil {
		return nil
	}

	if value.Equals(n.Value) {
		return n
	}

	if value.LessThan(n.Value) {
		return n.Left.Search(value)
	} else {
		return n.Right.Search(value)
	}
}

// Traverse prints the values in the binary tree in-order
func (n *BTree) Traverse() {
	if n != nil {
		n.Left.Traverse()
		n.Right.Traverse()
	}
}
