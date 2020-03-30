package distancetable

import "testing"

func TestDistanceTable(t *testing.T) {
	dt := NewDistanceTable()
	dt.Update(1, 2, 3)
	dt.NilDtUnit(0)

	if dt.DTUnitValue(0).Assigned() != false {
		t.Error("DT Nil Initialization is not working")
		t.FailNow()
	}

	if dt.DTUnitValue(1).Assigned() != true {
		t.Error("DT Update Assigmnet is not working")
		t.FailNow()
	}

	if dt.DTUnitValue(1).Weight() != 2 {
		t.Error("DT For Vertex Weight is incorrect")
		t.FailNow()
	}

	if dt.DTUnitValue(1).PreNode() != 3 {
		t.Error("DT For Vertex PreNode is incorrect")
		t.FailNow()
	}

	dt.Update(1, 5, 5)
	if dt.DTUnitValue(1).Weight() != 5 {
		t.Error("DT For Vertex Updated Weight is incorrect")
		t.FailNow()
	}

	if dt.DTUnitValue(1).PreNode() != 5 {
		t.Error("DT For Vertex Updated PreNode is incorrect")
		t.FailNow()
	}
}
