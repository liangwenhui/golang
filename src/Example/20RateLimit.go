package main

import (
	"fmt"
	"time"
)

//限流 通过 channel ticket 实现限流
func main() {
	//------每秒接受一个请求
	reqs := make(chan int, 5)
	//填充信道
	for i := 0; i < 5; i++ {
		reqs <- i
	}
	close(reqs)
	limter := time.Tick(time.Second)
	for req := range reqs {
		<-limter
		fmt.Print(req, " ")
	}
	fmt.Println()
	//-----------每秒接受3个请求
	reqs = make(chan int, 5)
	for i := 0; i < 5; i++ {
		reqs <- i
	}
	close(reqs)
	//限流器burstyLimter
	burstyLimter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimter <- time.Now()
	}
	go func() {
		for t := range time.Tick(time.Second) {
			burstyLimter <- t
		}
	}()
	for req := range reqs {
		<-burstyLimter
		fmt.Print(req)
	}

}
