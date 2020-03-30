package main

import (
	"fmt"

	"github.com/VimleshS/funWithGraph/algorithm"
	"github.com/VimleshS/funWithGraph/graphds/adjacencymatrix"
	"github.com/VimleshS/funWithGraph/graphds/adjacencyset"
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

	algorithm.TopologicalSortusingAdjacencySet(&a)

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
	algorithm.TopologicalSortusingAdjacencySet(&matrix)

	//SPW
	fmt.Println("SPUw")
	spm := adjacencymatrix.NewAdjacencyMatrix(8, false)
	spm.AddEdge(0, 1, 1)
	spm.AddEdge(1, 2, 1)
	spm.AddEdge(1, 3, 1)
	spm.AddEdge(2, 3, 1)
	spm.AddEdge(1, 4, 1)
	spm.AddEdge(3, 5, 1)
	spm.AddEdge(5, 4, 1)
	spm.AddEdge(3, 6, 1)
	spm.AddEdge(6, 7, 1)
	spm.AddEdge(0, 7, 1)
	algorithm.ShortestPathUnWeighted(&spm, 0, 5)
	algorithm.ShortestPathUnWeighted(&spm, 0, 6)
	algorithm.ShortestPathUnWeighted(&spm, 7, 4)

}
