package queue

import (
	"math/rand/v2"
	"testing"
)

// Test de implementación de interfaz Queue
func TestImplementsGraphInterface(t *testing.T) {
	var _ Queue[int] = &SliceQueue[int]{}
}

// Test de operación Empty
func TestEmpty(t *testing.T) {
	s := CreateSliceQueue[int]()

	if !s.Empty() {
		t.Errorf("Queue is not empty")
	}

	in := rand.IntN(1000)
	s.Enqueue(in)
	if s.Empty() {
		t.Errorf("Queue is empty")
	}
}

// Test de operaciones enqueue / dequeue / peek
func TestQueue(t *testing.T) {
	q := CreateSliceQueue[int]()

	// test : empty(create)
	if !q.Empty() {
		t.Errorf("Queue is not empty")
		return
	}

	slc := make([]int, 100000)
	for i := range slc {
		n := rand.IntN(1000)
		slc[i] = n
		q.Enqueue(n)
	}

	for i := range slc {
		if slc[i] != q.Peek() {
			t.Errorf("Peek operation failed")
			return
		}
		if slc[i] != q.Dequeue() {
			t.Errorf("Dequeue operation failed")
			return
		}
	}

	if !q.Empty() {
		t.Errorf("Queue is not empty after dequeue")
		return
	}
}
