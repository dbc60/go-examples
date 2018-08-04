package vending

import "github.com/dbc60/go-csp"

// VendChoc returns the next process and event (state) for VendChoc
func VendChoc(e csp.Event) (csp.Process, csp.Event) {
	if e == csp.Event(coin) {
		// This is a reset function
		reset := func(e csp.Event) (csp.Process, csp.Event) {
			if e == csp.Event(choc) {
				return VendChoc, csp.Event(coin)
			}
			return csp.Stop, csp.Bleep
		}
		return reset, csp.Event(choc)
	}
	// Broken process. It needed a coin, but got something else (a slug?).
	return csp.Stop, csp.Bleep
}
