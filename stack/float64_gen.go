// auto-generated

package stack

// Float64Stack is a stack of float64 values.
type Float64Stack struct {
	stack []float64
}

// Pop a value off the stack.
func (s *Float64Stack) Pop() float64 {
	if len(s.stack) == 0 {
		return 0
	}

	i := len(s.stack) - 1
	v := s.stack[i]
	s.stack = s.stack[:i]
	return v
}

// Peek at the top of the stack.
func (s *Float64Stack) Peek() float64 {
	if len(s.stack) == 0 {
		return 0
	}

	i := len(s.stack) - 1
	v := s.stack[i]
	return v
}

// Push a value onto the stack.
func (s *Float64Stack) Push(v float64) {
	s.stack = append(s.stack, v)
}

// Len returns the length of the stack.
func (s *Float64Stack) Len() int {
	return len(s.stack)
}

// Walk over every value in the stack.
func (s *Float64Stack) Walk(f func(float64)) {
	for _, v := range s.stack {
		f(v)
	}
}
