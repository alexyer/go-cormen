package bintree

import "fmt"

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
func Walk(t *Tree, ch chan *Tree) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t
	Walk(t.Right, ch)
}

// Launch Walk in a new goroutine and return read-only channel of values.
func Walker(t *Tree) <-chan *Tree {
	ch := make(chan *Tree)
	go func() {
		Walk(t, ch)
		close(ch)
	}()
	return ch
}

// Create new tree holding the values from 0 to n-1.
func NewRange(n int) *Tree {
	var t *Tree
	for i := 0; i < n; i++ {
		t = Insert(t, i)
	}
	return t
}

// Create new empty tree.
func New() *Tree {
	return &Tree{nil, 0, nil, nil}
}

// Insert new node into the tree.
func Insert(t *Tree, val int) *Tree {
	return InsertNode(t, t, val)
}

// Navigate through tree and insert node in proper place.
func InsertNode(t, p *Tree, val int) *Tree {
	if t == nil {
		return &Tree{nil, val, nil, nil}
	}
	if val < t.Value {
		t.Left = InsertNode(t.Left, t, val)
		t.Parent = p
		return t
	}
	t.Right = InsertNode(t.Right, t, val)
	t.Parent = p
	return t
}

// The procedure begins its search at the root and traces a simple path downward in the tree.
func Search(t *Tree, val int) *Tree {
	if t == nil || val == t.Value {
		return t
	}

	if val < t.Value {
		return Search(t.Left, val)
	} else {
		return Search(t.Right, val)
	}
}

// Return node with minimum value
func Min(t *Tree) *Tree {
	x := t

	for x.Left != nil {
		x = x.Left
	}

	return x
}

// Return node with maximum value
func Max(t *Tree) *Tree {
	x := t

	for x.Right != nil {
		x = x.Right
	}

	return x
}

// Return successor of the given node
func Successor(t *Tree) *Tree {
	if t.Right != nil {
		return Min(t.Right)
	}

	y := t.Parent
	for y != nil && t == y.Right {
		t = y
		y = y.Parent
	}

	return y
}

// Return predecessor of the given node
func Predecessor(t *Tree) *Tree {
	if t.Left != nil {
		return Max(t.Left)
	}

	y := t.Parent
	for y != nil && t == y.Left {
		t = y
		y = y.Parent
	}

	return y
}
