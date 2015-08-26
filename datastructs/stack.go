package datastructs

import "errors"

type stack struct {
	data []int
	top  int
	size int
}

func NewStack(size int) *stack {
	stack := new(stack)

	stack.top = -1
	stack.size = size
	stack.data = make([]int, size)

	return stack
}

func (s *stack) StackEmpty() bool {
	return s.top == -1
}

func (s *stack) Push(val int) error {
	if s.top+1 < s.size {
		s.top++
		s.data[s.top] = val
		return nil
	} else {
		return errors.New("Stack is full")
	}
}

func (s *stack) Pop() (int, error) {
	if s.StackEmpty() {
		return 0, errors.New("Stack is empty")
	}

	s.top--
	return s.data[s.top+1], nil
}
