package main

import (
	"fmt"

	"github.com/VimleshS/funWithGraph/graphds"
	"github.com/VimleshS/funWithGraph/graphds/adjacencymatrix"
	"github.com/VimleshS/funWithGraph/graphds/adjacencyset"
	"github.com/VimleshS/funWithGraph/queue"
)

func main() {
	//Implementation is based on hashMap hence result be different for sucessive runs
	a := adjacencyset.NewAdjacencySet(9, true)
	a.AddEdge(0, 1, 1)
	a.AddEdge(1, 2, 1)
	a.AddEdge(2, 7, 1)
	a.AddEdge(2, 4, 1)
	a.AddEdge(2, 3, 1)
	a.AddEdge(1, 5, 1)
	a.AddEdge(5, 6, 1)
	a.AddEdge(3, 6, 1)
	a.AddEdge(3, 4, 1)
	a.AddEdge(6, 8, 1)

	topologicalSortusingAdjacencySet(&a)

	matrix := adjacencymatrix.NewAdjacencyMatrix(9, true)
	matrix.AddEdge(0, 1, 1)
	matrix.AddEdge(1, 2, 1)
	matrix.AddEdge(2, 7, 1)
	matrix.AddEdge(2, 4, 1)
	matrix.AddEdge(2, 3, 1)
	matrix.AddEdge(1, 5, 1)
	matrix.AddEdge(5, 6, 1)
	matrix.AddEdge(3, 6, 1)
	matrix.AddEdge(3, 4, 1)
	matrix.AddEdge(6, 8, 1)
	topologicalSortusingAdjacencySet(&matrix)
}

func topologicalSortusingAdjacencySet(a graphds.IGraph) {
	topologicallySorted := []int{}
	q := queue.Queue{}

	indegreeMap := map[int]int{}
	for i := 0; i < a.NumVertices(); i++ {
		indegreeMap[i] = a.Indegree(i)
		if indegreeMap[i] == 0 {
			q.Push(i)
		}
	}

	for q.Len() > 0 {
		v := q.Pop()
		topologicallySorted = append(topologicallySorted, v)

		for _, _v := range a.AdjacentVertices(v) {
			indegreeMap[_v] = indegreeMap[_v] - 1
			if indegreeMap[_v] == 0 {
				q.Push(_v)
			}
		}
	}
	fmt.Println(topologicallySorted)
	//[0 1 2 5 3 7 4 6 8]
}
