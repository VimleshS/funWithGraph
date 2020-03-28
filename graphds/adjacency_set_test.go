package graphds

import (
	"reflect"
	"testing"
)

func TestAdjacencyCreate(t *testing.T) {
	const (
		size     = 3
		directed = true
	)
	a := NewAdjacencySet(size, directed)
	if len(*(a.store)) != size {
		t.Error("Not a valid Adjacency Set")
		t.Fail()
	}

	for _, s := range a.Store() {
		if s.Store() == nil {
			t.Error("Adjacency Sets internal set are not initailized")
		}
	}
	t.Log(a.directed)
}

func TestAdjacencySetIndegree(t *testing.T) {
	const (
		size     = 4
		directed = true
	)
	a := NewAdjacencySet(size, directed)
	a.AddEdge(0, 1)
	a.AddEdge(0, 2)
	a.AddEdge(2, 1)
	a.AddEdge(2, 3)

	testcases := []struct {
		valueUnderTest int
		expectedValue  int
		expectedSlice  []int
	}{
		{1, 2, []int{0, 2}},
		{3, 1, []int{2}},
		{0, 0, []int{}},
	}

	for _, _test := range testcases {
		indegree := a.Indegree(_test.valueUnderTest)

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
func TestAdjacencySetAdjacentVertices(t *testing.T) {
	const (
		size     = 4
		directed = true
	)
	a := NewAdjacencySet(size, directed)
	a.AddEdge(0, 1)
	a.AddEdge(0, 2)
	a.AddEdge(2, 1)
	a.AddEdge(2, 3)

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
