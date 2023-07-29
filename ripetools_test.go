package RIPEtools

import "testing"

func TestGetCountry(t *testing.T) {
	er, cn := GetCountry("8.8.8.8")
	if er != nil {
		t.Errorf("Get error: %s", er)
	} else {
		if cn != "US" {
			t.Errorf("Expected US but readed: %s", cn)
		}
	}
}
