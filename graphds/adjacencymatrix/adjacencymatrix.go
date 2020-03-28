package adjacencymatrix

//WeightInt a place holder for weight and if not nil means has a edge
type WeightInt struct {
	weight int

	//variable is notnil instead of null, so that defulat values remains not
	//assigned, since bool is bydefault false
	//null bool
	assigned bool
}

//Value Simulating assiging a nil array
func (w WeightInt) Value() interface{} {
	if w.assigned {
		return w.weight
	}
	return nil
}

//Assigned ...
func (w WeightInt) Assigned() bool {
	return w.assigned
}

//NewWeightInt not nil
func newWeightInt(v int) WeightInt {
	return WeightInt{v, true}
}

//AdjacencyMatrix ...
type AdjacencyMatrix struct {
	store       *([][]WeightInt)
	directed    bool
	numVertices int
}

//NewAdjacencyMatrix ...
func NewAdjacencyMatrix(size int, directed bool) AdjacencyMatrix {
	s := make([][]WeightInt, size)
	for i := 0; i < len(s); i++ {
		s[i] = make([]WeightInt, size)
	}
	g := AdjacencyMatrix{store: &s, directed: directed, numVertices: size}
	return g
}

//AddEdge ...
func (g *AdjacencyMatrix) AddEdge(v1, v2, weight int) {
	if v1 < 0 || v2 > len(*g.store) {
		panic("incorrect vertices gived to add")
	}
	(*(g.store))[v1][v2] = newWeightInt(weight)

	if g.directed == false {
		(*(g.store))[v2][v1] = newWeightInt(weight)
	}
}

func contains(a []WeightInt, v int) bool {
	for _, e := range a {
		if e.Value() == v {
			return true
		}
	}
	return false
}

//Indegree ...
func (g *AdjacencyMatrix) Indegree(v int) int {
	indegree := 0
	for _, rows := range *(g.store) {
		if rows[v].Assigned() {
			indegree++
		}
	}

	return indegree
}

//AdjacentVertices ...
func (g *AdjacencyMatrix) AdjacentVertices(v int) []int {
	adj := []int{}
	l := (*(g.store))[v]
	for i, v := range l {
		if v.Assigned() {
			adj = append(adj, i)
		}
	}
	return adj
}

//EdgeWeight ...
func (g *AdjacencyMatrix) EdgeWeight(v1, v2 int) int {
	w, ok := (*(g.store))[v1][v2].Value().(int)
	if ok {
		return w
	}
	return -1
}

//NumVertices ...
func (g *AdjacencyMatrix) NumVertices() int {
	return g.numVertices
}
