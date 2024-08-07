package stack

import (
	"math/rand/v2"
	"testing"
)

// Benchmark rendimiento SliceQueue
func BenchmarkSliceQueue(b *testing.B) {
	q := CreateSliceStack[int]()
	for i := 0; i < b.N; i++ {
		q.Push(rand.IntN(1000))
	}

	for !q.Empty() {
		q.Pop()
	}
}
