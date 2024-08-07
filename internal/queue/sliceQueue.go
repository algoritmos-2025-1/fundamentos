package queue

// Implementación de estructura Queue usando slices
type SliceQueue[T any] struct {
	// código
}

// Crea una estructura Queue
func CreateSliceQueue[T any]() *SliceQueue[T] {
	// código
	return nil
}

// Encola un elemento
func (q *SliceQueue[T]) Enqueue(item T) {
	// código
}

// Desencola un elemento
func (q *SliceQueue[T]) Dequeue() T {
	// código
	var x T
	return x
}

// Regresa el primer elemento
func (q *SliceQueue[T]) Peek() T {
	// código
	var x T
	return x
}

// Verifica si la cola es vacía
func (q *SliceQueue[T]) Empty() bool {
	// código
	return false
}
