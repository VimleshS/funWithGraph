package graphds

//IGraph Datastructure
type IGraph interface {
	AddEdge(v1, v2, weight int)
	Indegree(v int) int
	AdjacentVertices(v int) []int
	EdgeWeight(v1, v2 int) int
	NumVertices() int
}
