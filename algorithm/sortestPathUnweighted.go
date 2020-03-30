package algorithm

import (
	"fmt"

	"github.com/VimleshS/funWithGraph/graphds"
	"github.com/VimleshS/funWithGraph/graphds/distancetable"
	"github.com/VimleshS/funWithGraph/queue"
)

var dt distancetable.DistanceTable

func prepareDistanceTable(g graphds.IGraph, src, dst int) {
	dt = distancetable.NewDistanceTable()
	for i := 0; i < g.NumVertices(); i++ {
		dt.NilDtUnit(i)
	}

	//seed for start node
	dt.Update(src, dst, src)
	q := queue.Queue{}
	q.Push(src)

	for q.Len() > 0 {
		v := q.Pop()
		curDistance := dt.DTUnitValue(v).Weight()

		for _, adjVertices := range g.AdjacentVertices(v) {
			neighbourWeight := curDistance + 1
			if dt.DTUnitValue(adjVertices).Assigned() == false {
				dt.Update(adjVertices, neighbourWeight, v)
				q.Push(adjVertices)
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
	//useless code to simulate stack,
	_t := []int{}
	for i := len(r) - 1; i >= 0; i-- {
		_t = append(_t, r[i])
	}
	fmt.Printf("Shorted path unweighted from %d to %d is %v \n", src, dst, _t)

}

//ShortestPathUnWeighted ...
func ShortestPathUnWeighted(g graphds.IGraph, src, dst int) {
	prepareDistanceTable(g, src, dst)
}
