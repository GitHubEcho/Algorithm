package stack

import "errors"

//the size of the stack is allocated based on the slice
type Stack struct {
	data []interface{}
}

// Push adds x to the top of the stack.
func (s *Stack) Push(x interface{}) {
	s.data = append(s.data, x)
}

// Pop removes and returns the top element of the stack.
// It's a run-time error to call Pop on an empty stack.
func (s *Stack) Pop() (interface{}, error) {
	if s.Size() == 0 {
		return nil, errors.New("stack is empty")
	}
	i := len(s.data) - 1
	res := s.data[i]
	s.data[i] = nil // to avoid memory leak
	s.data = s.data[:i]
	return res, nil
}

// Size returns the number of elements in the stack.
func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) IsEmpty() bool {
	return len(s.data) == 0
}
