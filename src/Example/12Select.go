package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for i := 0; i < 3; i++ {
			c1 <- "c1"
			time.Sleep(time.Microsecond)
		}

	}()
	go func() {
		for i := 0; i < 3; i++ {
			c2 <- "c2"
			time.Sleep(time.Microsecond)
		}
	}()

	for i := 0; i < 1; i++ {
		select {
		case m1 := <-c1:
			fmt.Println("1cr1", m1)
		case m1 := <-c1:
			fmt.Println("1cr2", m1)
		case m1 := <-c2:
			fmt.Println("2cr1", m1)
		case m1 := <-c2:
			fmt.Println("2cr2", m1)
		}
	}

}
