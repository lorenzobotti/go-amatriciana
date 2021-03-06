package amatriciana

import (
	"testing"
)

func TestXyFromString(t *testing.T) {
	coord := XyFromString("e1")
	if coord != 0x04 {
		t.Fail()
	}
}

func TestXyToString(t *testing.T) {
	xy := Xy(0x70)
	if xy.String() != "a8" {
		t.Fail()
	}
}
