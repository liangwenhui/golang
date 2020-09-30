package main

import (
	"fmt"
	"time"
)

func thread1() {
	time.Sleep(1000)
	fmt.Println("t1")
}

func thread2() {
	fmt.Println("t2")
}

func main() {
	go thread1()
	go thread2()
	time.Sleep(2000)
}
