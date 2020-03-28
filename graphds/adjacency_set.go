package graphds

import (
	"github.com/VimleshS/funWithGraph/set"
)

//AdjacencySet ...
type AdjacencySet struct {
	//store is a pointer to array of set.Set pointers

	//  (*(a.store))[0] = set.NewSet()
	//  |----------| took address and then enumerated on whole to retrive set pointer
	//is a correct way to access, get the address of array of pointer and then
	//access it via index

	//*(a.store)[0] = set.NewSet() is a wrong bcoz we are not accessing the array
	// typecast error

	store *([]*set.Set)
	//a helper method Store is added to read store without fuss

	//Note: this defination is overkill it could simply be defined as
	//store []*set.Set
	directed bool

	//to hold num of vertices
	numVertices int
}

//NewAdjacencySet ...
func NewAdjacencySet(size int, directed bool) AdjacencySet {
	s := make([]*set.Set, size)
	//another way to achive
	// for i := range s {
	// 	s[i] = set.NewSet()
	// }

	a := AdjacencySet{store: &s, directed: directed, numVertices: size}
	for i := range *a.store {
		(*a.store)[i] = set.NewSet()
	}
	return a
}

//AddEdge ...
func (g *AdjacencySet) AddEdge(v1, v2, weight int) {
	if v1 < 0 || v2 > len(*(g.store))-1 {
		panic("wrong size")
	}
	(*g.store)[v1].AddElement(v2, weight)

	if g.directed == false {
		(*g.store)[v2].AddElement(v1, weight)
	}
}

//Indegree ...
func (g *AdjacencySet) Indegree(v int) int {
	if v < 0 || v > len(*(g.store))-1 {
		panic("wrong size")
	}
	t := []int{}
	for i, s := range *g.store {
		if s.Contain(v) {
			t = append(t, i)
		}
	}

	return len(t)
}

//AdjacentVertices ...
func (g *AdjacencySet) AdjacentVertices(v int) []int {
	if v < 0 || v > len(*(g.store))-1 {
		panic("wrong size")
	}
	return (*(g.store))[v].Elements()

}

//EdgeWeight ...
func (g *AdjacencySet) EdgeWeight(v1, v2 int) int {
	return (*(g.store))[v1].Weight(v2)
}

//Store ...
func (g *AdjacencySet) Store() []*set.Set {
	return *g.store
}

//NumVertices ...
func (g *AdjacencySet) NumVertices() int {
	return g.numVertices
}
