package main

import (
	"fmt"
	"time"
)

//select 加上超时等待操作
func main() {
	c := make(chan string, 1)
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c <- "go"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "go2"
	}()
	select {
	case msg := <-c:
		fmt.Println(msg)
	case time := <-time.After(time.Second):
		fmt.Println(time)
	}

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case time := <-time.After(3 * time.Second):
		fmt.Println(time)
	}

}
