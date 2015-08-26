package bintree

import "fmt"

// Binary tree with integer values.
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

func (t *Tree) String() string {
	return fmt.Sprintf("{Left node: %p, Value: %d, Right node: %p}", t.Left, t.Value, t.Right)
}

// Traverse a tree depth-first, sending each Value on a channel.
func Walk(t *Tree, ch chan int) {
	if t == nil {
		return
	}
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Launch Walk in a new goroutine and return read-only channel of values.
func Walker(t *Tree) <-chan int {
	ch := make(chan int)
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
	return &Tree{nil, 0, nil}
}

// Insert new node into the tree.
func Insert(t *Tree, val int) *Tree {
	if t == nil {
		return &Tree{nil, val, nil}
	}
	if val < t.Value {
		t.Left = Insert(t.Left, val)
		return t
	}
	t.Right = Insert(t.Right, val)
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
