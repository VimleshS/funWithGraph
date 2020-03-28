package set

//Set go impl
type Set struct {
	store map[int]bool
}

//NewSet ...
func NewSet() *Set {
	return &Set{
		store: map[int]bool{},
	}
}

//AddElement ...
func (s *Set) AddElement(v int) {
	s.store[v] = true
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

//Store a helper function to inspect store
func (s *Set) Store() map[int]bool {
	return s.store
}
