package set

//Set go impl
type Set struct {
	store map[int]int
}

//NewSet ...
func NewSet() *Set {
	return &Set{
		store: map[int]int{},
	}
}

//AddElement ...
func (s *Set) AddElement(v, weight int) {
	s.store[v] = weight
}

//Elements ...
func (s *Set) Elements() []int {
	elems := []int{}
	for e := range s.store {
		elems = append(elems, e)
	}
	return elems
}

//Contain ...
func (s *Set) Contain(v int) bool {
	if _, ok := s.store[v]; ok {
		return true
	}
	return false
}

func (s *Set) Weight(v int) int {
	if w, ok := s.store[v]; ok {
		return w
	}
	return -1
}

//Store a helper function to inspect store
func (s *Set) Store() map[int]int {
	return s.store
}
