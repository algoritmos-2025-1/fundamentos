package queue

type Queue[T any] interface {
	// Encola un elemento
	Enqueue(item T)

	// Desencola un elemento
	Dequeue() T

	// Regresa el primer elemento
	Peek() T

	// Verifica si la cola es vacía
	Empty() bool
}
