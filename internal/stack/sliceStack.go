package stack

// Implementación de estructura Stack usando slices
type SliceStack[T any] struct {
	items []T
}

// Crea una estructura Stack
func CreateSliceStack[T any]() *SliceStack[T] {
	// código
	return nil
}

// Inserta un elemento en la pila
func (s *SliceStack[T]) Push(item T) {
	// código
}

// Retira un elemento de la pila
func (s *SliceStack[T]) Pop() T {
	// código
	var v T
	return v
}

// Regresa el primer elemento
func (s *SliceStack[T]) Peek() T {
	// código
	var v T
	return v
}

// Verifica si la pila es vacía
func (s *SliceStack[T]) Empty() bool {
	// código
	return false
}
