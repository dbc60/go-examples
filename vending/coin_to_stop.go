package vending

import "github.com/dbc60/go-csp"

// CoinToStop accepts one coin, then stops
func CoinToStop(e csp.Event) (csp.Process, csp.Event) {
	if e == csp.Event(coin) {
		return csp.Stop, e
	}
	return csp.Stop, csp.Bleep
}

// NewCoinToStop uses prefix to return a Process that behaves like CoinToStop
func NewCoinToStop() csp.Process {
	return csp.Prefix(csp.Event(coin), csp.Stop)
}
