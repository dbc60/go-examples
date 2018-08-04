package vending

import (
	"github.com/dbc60/go-csp"
	"testing"
)

func TestVendChoc(t *testing.T) {
	m, e := VendChoc, csp.Event(coin)
	for i := 0; i < 4; i++ {
		if e != csp.Event(coin) {
			t.Errorf("Deposit: unexpected event %v", e)
		} else {
			t.Logf("Deposit: %v", Event(e))
		}
		m, e = m(e)
		if e != csp.Event(choc) {
			t.Errorf("Vend: unexpected event %v", e)
		} else {
			t.Logf("Vend: %v", Event(e))
		}
		m, e = m(e)
	}
}
