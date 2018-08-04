package main

import (
	"fmt"
	"time"
)

func f(left, right chan int) {
	left <- 1 + <-right
}

func main() {
	const n = 100000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost

	// Let's time this thing
	start := time.Now()

	/* Make a chain of channels and gophers from left-to-right, calling each
	 * gopher in the chain. Note that the new gopher's right channel becomes
	 * the next gopher's left channel.
	 *
	 * None of the gophers have received any data on the right, so they will
	 * each wait on their right channel.
	 */
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		left = right
	}

	// Kick off a function to send data on the rightmost channel.
	go func(c chan int) { c <- 1 }(right)

	// Wait for data on the leftmost channel and get the elapsed time.
	value := <-leftmost
	elapsed := time.Since(start)

	// Display the results
	fmt.Println(value)
	fmt.Println(elapsed)
}
