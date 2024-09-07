package main

import (
	"fmt"

	"github.com/algoritmos-2025-1/fundamentos/internal/collection"
	"github.com/algoritmos-2025-1/fundamentos/internal/stack"
	"github.com/algoritmos-2025-1/fundamentos/pkg/algorithms"
	"github.com/algoritmos-2025-1/fundamentos/pkg/graph"
)

func main() {
	// q := queue.CreateSliceQueue[int]()
	// q.Enqueue(42)
	// fmt.Println(q.Dequeue())
	
	s := stack.CreateSliceStack[int]()
	s.Push(0)

	c := collection.Collection[int](s)

	g := graph.CreateAdjacencyListGraph(5).(*graph.AdjacencyListGraph)
	g.AddEdge(0, 1)
	g.AddEdge(0, 2)
	g.AddEdge(0, 3)
	g.AddEdge(0, 4)
	g.AddEdge(1, 4)
	g.AddEdge(1, 2)
	g.AddEdge(4, 2)
	g.Print()

	p, t := algorithms.GraphSearch(g, 0, c)
	fmt.Println(p)
	fmt.Println(t)
}