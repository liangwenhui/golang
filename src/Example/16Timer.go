package main

import (
	"fmt"
	"time"
)

//计时器
func main() {
	timer := time.NewTimer(time.Second)

	go func() {
		time := <-timer.C
		fmt.Println(time)
	}()
	timer.Stop()
	time.Sleep(2 * time.Second)

}
