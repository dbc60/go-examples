package vending

import "testing"

func TestVendEventString(t *testing.T) {
	if "stop" != stop.String() {
		t.Errorf("stop.String returned %s", stop.String())
	}
	if "coin" != coin.String() {
		t.Errorf("coin.String returned %s", coin.String())
	}
	if "toffee" != toffee.String() {
		t.Errorf("coin.String returned %s", toffee.String())
	}
	if "chocolate" != choc.String() {
		t.Errorf("coin.String returned %s", choc.String())
	}
}
