package main

import (
	"fmt"
	"time"
)

/*
go 中线程通信 通过channel
*/
func get(c chan string) {
	for {
		s := <-c
		fmt.Println("get1", s)
	}

}
func get2(c chan string) {
	for {
		s := <-c
		fmt.Println("get2", s)
	}

}

func main() {

	ch := make(chan string)
	go get(ch)
	go get2(ch)
	fmt.Println(ch)
	for {
		ch <- "hello!"
		time.Sleep(time.Second)
	}

}
