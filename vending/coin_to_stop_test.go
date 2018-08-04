package vending

import (
	"github.com/dbc60/go-csp"
	"testing"
)

// TestCoinToStop tests the state transitions for CoinToStop
func TestCoinToStop(t *testing.T) {
	m, e := CoinToStop, csp.Event(coin)
	m, e = m(e)
	if e != csp.Event(coin) {
		t.Errorf("Unexpected event %v", e)
	}
	m, e = m(e)
	if e != csp.Bleep {
		t.Errorf("Unexpected event %v", e)
	}
}

func TestNewCoinToStop(t *testing.T) {
	m := NewCoinToStop()
	e := csp.Event(coin)
	m, e = m(e)
	if e != csp.Event(coin) {
		t.Errorf("Unexpected event %v", e)
	}
	m, e = m(e)
	if e != csp.Bleep {
		t.Errorf("Unexpected event %v", e)
	}
}
