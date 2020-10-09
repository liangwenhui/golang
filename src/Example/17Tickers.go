package main

import (
	"fmt"
	"time"
)

//循环计时器
func main() {

	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case t := <-ticker.C:
				fmt.Println(t)
			case <-done:
				return
			}
		}
	}()
	time.Sleep(3 * time.Second)
	ticker.Stop()
	done <- true

}
