package gominpriorityqueue

// This example demonstrates a priority queue built using the heap interface.

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	vertex           int // The value of the item; arbitrary.
	weightedPriority int // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

//NewItem is new queue element
func NewItem(vertex, weightedPriority int) *Item {
	return &Item{
		vertex:           vertex,
		weightedPriority: weightedPriority,
	}
}

// Vertex ...
func (i *Item) Vertex() int {
	return i.vertex
}

//WeightPriority ...
func (i *Item) WeightPriority() int {
	return i.weightedPriority
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].weightedPriority < pq[j].weightedPriority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

//Push ...
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

//Pop ...
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, vertex int, weightedpriority int) {
	item.vertex = vertex
	item.weightedPriority = weightedpriority
	heap.Fix(pq, item.index)
}
