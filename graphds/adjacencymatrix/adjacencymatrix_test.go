package adjacencymatrix

import (
	"reflect"
	"testing"
)

func TestAdjacencymatrix(t *testing.T) {
	const (
		size     = 4
		directed = true
	)
	g := NewAdjacencyMatrix(size, directed)

	if len(*g.store) != size {
		t.Error("NewAdjacencyMatrix did not formed a correct matric level 1")
		t.FailNow()
	}
	for _, cols := range *g.store {
		if len(cols) != size {
			t.Error("NewAdjacencyMatrix did not formed a correct matric level 2")
			t.FailNow()
		}
	}
}

func TestAdjacencyMatrixWeight(t *testing.T) {
	const (
		size     = 4
		directed = true
	)
	g := NewAdjacencyMatrix(size, directed)

	if len(*g.store) != size {
		t.Error("NewAdjacencyMatrix did not formed a correct matric level 1")
		t.FailNow()
	}
	g.AddEdge(0, 1, 3)
	g.AddEdge(2, 1, 6)

	testcases := []struct {
		testFromVertex int
		testToVertex   int
		expectedValue  int
	}{
		{0, 1, 3},
		{2, 1, 6},
		{2, 2, -1},
	}

	for _, _test := range testcases {
		w := g.EdgeWeight(_test.testFromVertex, _test.testToVertex)
		if w != _test.expectedValue {
			t.Errorf("Failed for %d %d expected %d, instead got %d",
				_test.testFromVertex, _test.testToVertex, _test.expectedValue, w)
			t.FailNow()
		}
	}
}
func TestAdjacencySetWeigthUndirected(t *testing.T) {
	const (
		size     = 4
		directed = false
	)
	a := NewAdjacencyMatrix(size, directed)
	a.AddEdge(0, 1, 3)
	a.AddEdge(0, 2, 4)
	a.AddEdge(2, 1, 6)
	a.AddEdge(2, 3, 9)

	testcases := []struct {
		fromVertex     int
		toVertex       int
		expectedWeight int
	}{
		{0, 1, 3},
		{2, 3, 9},
		{3, 2, 9},
		{1, 0, 3},
	}

	for _, _test := range testcases {
		w := a.EdgeWeight(_test.fromVertex, _test.toVertex)

		if w != _test.expectedWeight {
			t.Errorf("From %d to %d Expected weight is %d instead calcuted as %d",
				_test.fromVertex, _test.toVertex, _test.expectedWeight, w)
			t.FailNow()
		}

	}
}

func TestAdjacencyMatrixIndegree(t *testing.T) {
	const (
		size     = 4
		directed = true
	)
	a := NewAdjacencyMatrix(size, directed)
	a.AddEdge(0, 1, 1)
	a.AddEdge(0, 2, 1)
	a.AddEdge(2, 1, 1)
	a.AddEdge(2, 3, 1)

	testcases := []struct {
		valueUnderTest int
		expectedValue  int
	}{
		{1, 2},
		{3, 1},
		{0, 0},
	}

	//debug
	// for _, rows := range *a.store {
	// 	t.Log(rows)
	// }

	for _, _test := range testcases {
		indegree := a.Indegree(_test.valueUnderTest)

		if indegree != _test.expectedValue {
			t.Errorf("For %d Expected Indegree is %d instead calcuted as %v",
				_test.valueUnderTest, _test.expectedValue, indegree)
			t.FailNow()
		}
	}
}

func TestAdjacencySetAdjacentVertices(t *testing.T) {
	const (
		size     = 4
		directed = true
	)
	a := NewAdjacencyMatrix(size, directed)
	a.AddEdge(0, 1, 1)
	a.AddEdge(0, 2, 1)
	a.AddEdge(2, 1, 1)
	a.AddEdge(2, 3, 1)

	testcases := []struct {
		valueUnderTest int
		expectedValue  int
		expectedSlice  []int
	}{
		{1, 0, []int{}},
		{2, 2, []int{1, 3}},
		{0, 2, []int{1, 2}},
	}

	for _, _test := range testcases {
		indegree := a.AdjacentVertices(_test.valueUnderTest)

		if len(indegree) != _test.expectedValue {
			t.Errorf("For %d Expected Indegree is %d instead calcuted as %v",
				_test.valueUnderTest, _test.expectedValue, len(indegree))
			t.FailNow()
		}

		if reflect.DeepEqual(indegree, _test.expectedSlice) == false {
			t.Errorf("For %d Expected Indegree is %v instead calcuted as %v",
				_test.valueUnderTest, _test.expectedSlice, indegree)
			t.FailNow()
		}
	}
}
