package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go func() {
		i := 0
		for i < 10 {
			s := fmt.Sprint("hello", i)
			i++
			c <- s
		}
	}()
	go func() {
		for {
			s := <-c
			fmt.Println(s)
		}
	}()
	time.Sleep(time.Second)
}
