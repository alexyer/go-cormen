package bst

import (
	"fmt"

	"github.com/alexyer/go-cormen/datastructs/trees"
)

// Binary tree with integer values.
type Tree struct {
	Left   *Tree
	Value  int
	Right  *Tree
	Parent *Tree
}

func (t *Tree) String() string {
	return fmt.Sprintf("{Left node: %p, Value: %d, Right node: %p, Parent Node: %p, Self: %p}", t.Left, t.Value, t.Right, t.Parent, t)
}

// Traverse a tree depth-first, sending each Value on a channel.
func (t *Tree) Walk(ch chan *Tree) {
	if t == nil {
		return
	}
	t.Left.Walk(ch)
	ch <- t
	t.Right.Walk(ch)
}

// Launch Walk in a new goroutine and return read-only channel of values.
func (t *Tree) Walker() <-chan *Tree {
	ch := make(chan *Tree)
	go func() {
		t.Walk(ch)
		close(ch)
	}()
	return ch
}

// Create new empty tree.
func New() *Tree {
	return &Tree{nil, 0, nil, nil}
}

// Insert new node into the tree.
func (t *Tree) Insert(val int) trees.Tree {
	var y *Tree = nil
	x := t

	z := New()
	z.Value = val

	for x != nil {
		y = x
		if z.Value < x.Value {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.Parent = y
	if z.Value < y.Value {
		y.Left = z
	} else {
		y.Right = z
	}

	return z
}

// The procedure begins its search at the root and traces a simple path downward in the tree.
func (t *Tree) Find(val int) trees.Tree {
	if t == nil || val == t.Value {
		return t
	}

	if val < t.Value {
		return t.Left.Find(val)
	} else {
		return t.Right.Find(val)
	}
}

// Return node with minimum value
func (t *Tree) Min() *Tree {
	x := t

	for x.Left != nil {
		x = x.Left
	}

	return x
}

// Return node with maximum value
func (t *Tree) Max() *Tree {
	x := t

	for x.Right != nil {
		x = x.Right
	}

	return x
}

// Return successor of the given node
func (t *Tree) Successor() *Tree {
	if t.Right != nil {
		return t.Right.Min()
	}

	y := t.Parent
	for y != nil && t == y.Right {
		t = y
		y = y.Parent
	}

	return y
}

// Return predecessor of the given node
func (t *Tree) Predecessor() *Tree {
	if t.Left != nil {
		return t.Left.Max()
	}

	y := t.Parent
	for y != nil && t == y.Left {
		t = y
		y = y.Parent
	}

	return y
}

// Replace one subtree as a child of its parent with another subtree.
func Transplant(t1, t2 *Tree) *Tree {
	if t1 == t1.Parent.Left {
		t1.Parent.Left = t2
	} else {
		t1.Parent.Right = t2
	}

	if t2 != nil {
		t2.Parent = t1.Parent
	}

	return t1
}

// Delete node from the tree.
func (t *Tree) Delete() {
	switch {
	case t.Left == nil:
		Transplant(t, t.Right)
	case t.Right == nil:
		Transplant(t, t.Left)
	default:
		y := t.Right.Min()
		if y.Parent != t {
			Transplant(y, y.Right)
			y.Right = t.Right
			y.Right.Parent = y
		}

		Transplant(t, y)
		y.Left = t.Left
		y.Left.Parent = y
	}
}
