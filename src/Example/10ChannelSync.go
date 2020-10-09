package main

import (
	"fmt"
	"time"
)

//channel的通信是同步的，发送等待接受，接受等待发送
func main() {
	c := make(chan bool)
	go func() {
		fmt.Println("working...")
		time.Sleep(3 * time.Second)
		fmt.Println("done")
		c <- true
	}()

	<-c

}
