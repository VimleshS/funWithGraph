package djikstra

import (
	"container/heap"
	"fmt"

	"github.com/VimleshS/funWithGraph/graphds"
	"github.com/VimleshS/funWithGraph/graphds/distancetable"
	minPq "github.com/VimleshS/funWithGraph/graphds/gominpriorityqueue"
)

//var dt distancetable.DistanceTable

func prepareDistanceTable(g graphds.IGraph, src, dst int) {
	dt := distancetable.NewDistanceTable()
	for i := 0; i < g.NumVertices(); i++ {
		dt.NilDtUnit(i)
	}

	//seed for start node
	dt.Update(src, 0, src)

	minPriorityQ := minPq.PriorityQueue{}
	heap.Push(&minPriorityQ, minPq.NewItem(src, 0))

	for minPriorityQ.Len() > 0 {
		qItem := heap.Pop(&minPriorityQ).(*minPq.Item)
		curDistance := dt.DTUnitValue(qItem.Vertex()).Weight()

		for _, adjVertices := range g.AdjacentVertices(qItem.Vertex()) {
			distance := curDistance + g.EdgeWeight(qItem.Vertex(), adjVertices)
			neighbourWeight := dt.DTUnitValue(adjVertices).Weight()
			if !dt.DTUnitValue(adjVertices).Assigned() || neighbourWeight > distance {
				dt.Update(adjVertices, distance, qItem.Vertex())
				heap.Push(&minPriorityQ, minPq.NewItem(adjVertices, distance))
			}
		}
	}
	backtrack(dt, src, dst)
}

func backtrack(dt distancetable.DistanceTable, src, dst int) {
	r := []int{}
	r = append(r, dst)
	pN := dst
	for dt.DTUnitValue(pN).Assigned() && pN != src {
		pN = dt.DTUnitValue(pN).PreNode()
		r = append(r, pN)
	}

	if dt.DTUnitValue(pN).Assigned() == false {
		fmt.Println("There does not exist a valid path")
	} else {
		//useless code to simulate stack,
		_t := []int{}
		for i := len(r) - 1; i >= 0; i-- {
			_t = append(_t, r[i])
		}
		fmt.Printf("Shorted path unweighted from %d to %d is %v \n", src, dst, _t)
	}

}

//ShortestPath ...
func ShortestPath(g graphds.IGraph, src, dst int) {
	prepareDistanceTable(g, src, dst)
}
