package RIPEtools

import "testing"

func TestGetCountry(t *testing.T) {
	var rd RIPEd
	ergd := rd.GetData("8.8.8.8")
	if ergd != nil {
		t.Errorf("GetData error: %s", ergd)
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
