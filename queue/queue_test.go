package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {
	q := Queue{}
	q.Push(2)
	q.Push(1)
	q.Push(3)

	a := []int{}
	for q.Len() > 0 {
		a = append(a, q.Pop())
	}
	if reflect.DeepEqual(a, []int{2, 1, 3}) == false {
		t.Error("Queue implementaion is incorrect")
		t.Fail()
	}
}
