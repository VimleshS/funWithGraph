package queue

//Queue naive queue of int
type Queue struct {
	store []int
}

//Push stores int to a array
func (q *Queue) Push(v int) {
	q.store = append(q.store, v)
}

//Pop ...
func (q *Queue) Pop() int {
	var _t int
	_t,q.store = q.store[0],q.store[1:]	 
	return _t
}

//Len ..usedas EOF
func (q *Queue) Len() int {
	return len(q.store)
}