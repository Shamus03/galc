#!/bin/bash

type_name="$1"
: ${type_name:?}
if [[ "$type_name" =~ " " ]]; then
  echo "type name must not contain spaces" &>2
  exit 1
fi

struct_name="${type_name^}Stack"

cat > ${type_name}_gen.go << EOF
// auto-generated

package stack

// ${struct_name} is a stack of ${type_name} values.
type ${struct_name} struct {
	stack []${type_name}
}

// Pop a value off the stack.
func (s *${struct_name}) Pop() ${type_name} {
	if len(s.stack) == 0 {
		return 0
	}

	i := len(s.stack) - 1
	v := s.stack[i]
	s.stack = s.stack[:i]
	return v
}

// Peek at the top of the stack.
func (s *${struct_name}) Peek() ${type_name} {
	if len(s.stack) == 0 {
		return 0
	}

	i := len(s.stack) - 1
	v := s.stack[i]
	return v
}

// Push a value onto the stack.
func (s *${struct_name}) Push(v ${type_name}) {
	s.stack = append(s.stack, v)
}

// Len returns the length of the stack.
func (s *${struct_name}) Len() int {
	return len(s.stack)
}

// Walk over every value in the stack.
func (s *${struct_name}) Walk(f func(${type_name})) {
	for _, v := range s.stack {
		f(v)
	}
}

// Rotate the stack (move the top value to the bottom of the stack)
func (s *${struct_name}) Rotate() {
	if len(s.stack) == 0 {
		return
	}

	i := 0
	v := s.stack[i]

	s.stack = append(s.stack[i+1:], v)
}

// Unrotate the stack (move the bottom value to the top of the stack)
func (s *${struct_name}) Unrotate() {
	if len(s.stack) == 0 {
		return
	}

	i := len(s.stack) - 1
	v := s.stack[i]

	s.stack = append([]${type_name}{v}, s.stack[:i]...)
}
EOF

exit 0
