package stack_test

import (
	"testing"

	"github.com/shamus03/galc/stack"
	"github.com/stretchr/testify/assert"
)

func Test_Stack_Push(t *testing.T) {
	var s stack.Float64Stack
	s.Push(1)
	assert.Equal(t, 1, s.Len())

	s.Push(2)
	assert.Equal(t, 2, s.Len())
}

func Test_Stack_Pop(t *testing.T) {
	var s stack.Float64Stack
	s.Push(1)
	s.Push(3)
	s.Push(2)

	assert.Equal(t, 2.0, s.Pop())
	assert.Equal(t, 3.0, s.Pop())
	assert.Equal(t, 1.0, s.Pop())
	assert.Equal(t, 0.0, s.Pop())
}

func Test_Stack_Walk(t *testing.T) {
	var s stack.Float64Stack
	s.Push(1)
	s.Push(3)
	s.Push(2)

	var vals []float64
	s.Walk(func(v float64) {
		vals = append(vals, v)
	})

	assert.Equal(t, []float64{1, 3, 2}, vals)
}

func Test_Stack_RotateUnrotate(t *testing.T) {
	var s stack.Float64Stack
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Rotate()
	assert.Equal(t, 1.0, s.Peek())

	s.Unrotate()
	assert.Equal(t, 3.0, s.Peek())

	assert.Equal(t, 3, s.Len())
}
