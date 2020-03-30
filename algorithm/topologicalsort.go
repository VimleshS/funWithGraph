package algorithm

import (
	"fmt"

	"github.com/VimleshS/funWithGraph/graphds"
	"github.com/VimleshS/funWithGraph/queue"
)

//TopologicalSortusingAdjacencySet ...
func TopologicalSortusingAdjacencySet(a graphds.IGraph) {
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
