package distancetable

//DTUnit to hold a value corresponding to a vertex
type DTUnit struct {
	weight   int
	prevNode int
	assigned bool
}

//Assigned to denote was this node value assigned
func (dt *DTUnit) Assigned() bool {
	return dt.assigned
}

//PreNode ...
func (dt *DTUnit) PreNode() int {
	return dt.prevNode
}

//Weight ...
func (dt *DTUnit) Weight() int {
	return dt.weight
}

//Update ...
func (dt *DTUnit) Update(w, pn int) {
	dt.assigned = true
	dt.weight = w
	dt.prevNode = pn
}

//DistanceTable ...
type DistanceTable struct {
	store map[int]*DTUnit
}

//NewDistanceTable ...
func NewDistanceTable() DistanceTable {
	return DistanceTable{
		store: make(map[int]*DTUnit),
	}
}

//Update to fill in distance table
func (d DistanceTable) Update(n, w, pn int) {
	if dte, ok := d.store[n]; ok {
		dte.Update(w, pn)
	} else {
		dte := DTUnit{}
		dte.Update(w, pn)
		d.store[n] = &dte
	}
}

//NilDtUnit to fill in distance table
func (d DistanceTable) NilDtUnit(n int) {
	d.store[n] = &DTUnit{}
}

//DTUnitValue to fill in distance table
func (d DistanceTable) DTUnitValue(n int) *DTUnit {
	return d.store[n]
}
