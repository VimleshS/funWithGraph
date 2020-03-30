package gominpriorityqueue

import (
	"container/heap"
	"testing"
)

func TestMinQueue(t *testing.T) {
	// Create a priority queue, put the items in it, and
	// establish the priority queue (heap) invariants.
	//pq := make(PriorityQueue, len(items))
	pq := PriorityQueue{}

	// Insert a new item and then modify its priority.
	item := &Item{
		vertex:           200,
		weightedPriority: 10,
	}
	heap.Push(&pq, item)
	pq.update(item, item.vertex, 5)
	heap.Push(&pq, &Item{vertex: 100, weightedPriority: 1})
	heap.Push(&pq, &Item{vertex: 66, weightedPriority: 6})
	heap.Push(&pq, &Item{vertex: 3, weightedPriority: 3})
	if pq.Len() != 4 {
		t.Error("Elements did mnot queued")
		t.FailNow()
	}

	wt := heap.Pop(&pq).(*Item)
	if wt.Vertex() != 100 {
		t.Errorf("Error in dequeueing expecting %d got %v \n", 100, wt)
		t.FailNow()
	}

	//Take the items out; they arrive in decreasing priority order.
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		t.Logf("%d  Weight %d Index %d \n", item.weightedPriority, item.vertex, item.index)
	}

}
