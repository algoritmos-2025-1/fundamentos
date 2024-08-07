package queue

import (
	"math/rand/v2"
	"testing"
)

// Benchmark rendimiento SliceQueue
func BenchmarkSliceQueue(b *testing.B) {
	q := CreateSliceQueue[int]()
	for i := 0; i < b.N; i++ {
		q.Enqueue(rand.IntN(1000))
	}

	for !q.Empty() {
		q.Dequeue()
	}
}
