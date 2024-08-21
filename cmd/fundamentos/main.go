package main

import (
	"fmt"

	"github.com/algoritmos-2025-1/fundamentos/internal/collection"
	"github.com/algoritmos-2025-1/fundamentos/internal/queue"
	"github.com/algoritmos-2025-1/fundamentos/pkg/algorithms"
	"github.com/algoritmos-2025-1/fundamentos/pkg/graph"
)

func main() {
	q := queue.CreateSliceQueue[int]()
	q.Enqueue(42)
	fmt.Println(q.Dequeue())

	c := collection.Collection[int](q)
	println(c)

	g := graph.CreateAdjacencyListGraph(12).(*graph.AdjacencyListGraph)
	g.AddEdge(0, 11)
	g.AddEdge(2, 11)
	g.AddEdge(3, 11)
	g.AddEdge(1, 10)
	g.AddEdge(0, 1)
	g.Print()

	println(g.Order())
	println(g.Size())

	p, t := algorithms.GraphSearch(g, 0, c)
	fmt.Println(p)
	fmt.Println(t)
}
