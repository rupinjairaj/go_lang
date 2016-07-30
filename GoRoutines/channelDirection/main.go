/*
* when channels are used as function parameters we can specifiy the direction
* of these channels i.e., if a channel is only meant to send or receive values.
* Makes it more type safe
 */

package main

import "fmt"

// This function only accepts values to send on the 'pings' channel
// would throw a compile time error if we try to receive from the 'pings' channel
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 'pings' is set to only receive values
// 'pongs' is set to only send values
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
