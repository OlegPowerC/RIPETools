package RIPEtools

import "testing"

func TestGetCountry(t *testing.T) {

	rd, rderr := NewRIPEreq("8.8.8.8")
	if rderr != nil {
		t.Errorf("Get error: %s", rderr)
	}
	er, cn := rd.GetCountry()
	if er != nil {
		t.Errorf("Get error: %s", er)
	} else {
		if cn != "US" {
			t.Errorf("Expected US but readed: %s", cn)
		}
	}
}
