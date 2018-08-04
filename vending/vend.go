package vending

import "github.com/dbc60/go-csp"

// Event is a csp.Event, used primarily for the String() thing
type Event csp.Event

const (
	stop = Event(csp.Bleep + iota)
	coin
	toffee
	choc
)

func (e Event) String() string {
	switch e {
	case stop:
		return "stop"
	case coin:
		return "coin"
	case toffee:
		return "toffee"
	case choc:
		return "chocolate"
	default:
		return "unknown vending event"
	}
}
