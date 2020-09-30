package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//读写锁
var rwlock sync.RWMutex

//countdownlauth
var wait sync.WaitGroup

//线程安全map
var currMap = sync.Map{}

func st() {
	currMap.Store("k", "v")
	load, ok := currMap.Load("k")
	fmt.Println(load, ok)
}

//原子操作
func at() {
	var num int32 = 0
	atomic.AddInt32(&num, 200)

}
func main() {
	rwlock.Lock()
	rwlock.RLock()
	wait.Add(2)
	wait.Done()
	wait.Done()
	wait.Wait()

}
