package main

import "fmt"

func main() {

	messages := make(chan string, 2)
	messages <- "buffered"
	messages <- "channel"
	// has to take 'buffered' from the other end else no more can be added to this channel
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	// can be added only if one of the other two which was buffered into the channel is removed and the pipe is open for more incoming data
	messages <- "hello"

	fmt.Println(<-messages)
}
