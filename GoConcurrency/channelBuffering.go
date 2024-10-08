package main

import "fmt"

func channelBuffering() {
	// make(chan string,n int) => "n int" mean n time buffer of channel.So we can receive n time buffer of channel without receiver.
	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
