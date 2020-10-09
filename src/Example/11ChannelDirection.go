package main

import "fmt"

func main() {
	ping := func(pings chan string, msg string) {
		pings <- msg
	}

	pong := func(pings chan string, pongs chan string) {
		pongs <- <-pings
	}
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "hello")
	pong(pings, pongs)
	fmt.Println(<-pongs)

}
