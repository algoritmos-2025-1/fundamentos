package stack

import (
	"math/rand/v2"
	"testing"
)

// Test de implementación de interfaz Stack
func TestImplementsGraphInterface(t *testing.T) {
	var _ Stack[int] = &SliceStack[int]{}
}

// Test de operación Empty
func TestEmpty(t *testing.T) {
	s := CreateSliceStack[int]()

	// test: push(empty)
	if !s.Empty() {
		t.Errorf("Stack is not empty")
	}

	in := rand.IntN(1000)
	s.Push(in)
	if s.Empty() {
		t.Errorf("Stack is empty")
	}
}

// Test de operaciones push / pop / peek
func TestStack(t *testing.T) {
	s := CreateSliceStack[int]()

	slc := make([]int, 100000)
	for i := range slc {
		n := rand.IntN(1000)
		slc[i] = n
		s.Push(n)
	}

	for i := len(slc) - 1; i > -1; i-- {
		if slc[i] != s.Peek() {
			t.Errorf("Peek operation failed")
			return
		}

		if slc[i] != s.Pop() {
			t.Errorf("Pop operation failed")
			return
		}
	}

	if !s.Empty() {
		t.Errorf("Stack is not empty after pop")
	}
}
