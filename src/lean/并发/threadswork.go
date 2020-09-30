package main

import (
	"fmt"
	"sync"
	"time"
)

/**
缓存为1的channel可以作为同步锁
*/
var lock = make(chan int, 1)
var sylock sync.Mutex
var moA int = 1000

/**
线程不安全
*/
func subMoa(i int) {
	time.Sleep(2 * time.Second)
	moA -= i
}

func subMoaLock(i int) {
	sylock.Lock()
	time.Sleep(2 * time.Second)
	moA -= i
	sylock.Unlock()
}
func getMoa() int {
	sylock.Lock()
	res := moA
	sylock.Unlock()
	return res
}
func main() {
	m2()

}

func m2() {
	go func() {
		if getMoa() > 200 {
			subMoaLock(200)
			fmt.Println("200扣款成功,剩余:", moA)
		}
	}()
	go func() {
		if getMoa() > 900 {
			subMoaLock(900)
			fmt.Println("900扣款成功,剩余:", moA)
		}
	}()
	time.Sleep(4 * time.Second)
	fmt.Println(moA)
}
func m1() {
	go func() {
		lock <- 10
		if moA > 200 {
			subMoa(200)
			fmt.Println("200扣款成功,剩余:", moA)
		}
		<-lock
	}()
	go func() {
		lock <- 10
		if moA > 900 {
			subMoa(900)
			fmt.Println("900扣款成功,剩余:", moA)
		}
		<-lock
	}()
	time.Sleep(4 * time.Second)
	fmt.Println(moA)
}
