package main

import (
	"fmt"

	"github.com/algoritmos-2025-1/fundamentos/internal/queue"
)

func main() {
	q := queue.CreateSliceQueue[int]()
	q.Enqueue(42)
	fmt.Println(q.Dequeue())
}
