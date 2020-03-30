package main

import (
	"fmt"

	"github.com/VimleshS/funWithGraph/algorithm/shortestpath/djikstra"
	"github.com/VimleshS/funWithGraph/graphds/adjacencymatrix"
)

func main() {
	/*
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
		spm := adjacencymatrix.NewAdjacencyMatrix(8, true)
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

		unweighted.ShortestPathUnWeighted(&spm, 0, 5)
		unweighted.ShortestPathUnWeighted(&spm, 0, 6)
		unweighted.ShortestPathUnWeighted(&spm, 7, 4)
	*/

	fmt.Println("Djikshtra ...")
	spd := adjacencymatrix.NewAdjacencyMatrix(8, true)
	spd.AddEdge(0, 1, 1)
	spd.AddEdge(1, 2, 2)
	spd.AddEdge(1, 3, 6)
	spd.AddEdge(2, 3, 2)
	spd.AddEdge(1, 4, 3)
	spd.AddEdge(3, 5, 1)
	spd.AddEdge(5, 4, 5)
	spd.AddEdge(3, 6, 1)
	spd.AddEdge(6, 7, 1)
	spd.AddEdge(0, 7, 8)

	djikstra.ShortestPath(&spd, 0, 6)
	djikstra.ShortestPath(&spd, 4, 7)
	djikstra.ShortestPath(&spd, 7, 0)

}
