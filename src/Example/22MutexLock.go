package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lock sync.Mutex
	var wg sync.WaitGroup
	pint := func(s string){

		fmt.Print(s)

	}

	go func() {
		wg.Add(1)
		for i:=0;i<10;i++ {
			lock.Lock()
			pint("A")
			lock.Unlock()
		}
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		for i:=0;i<10;i++ {
			lock.Lock()
			pint("B")
			lock.Unlock()
		}
		wg.Done()
	}()
	time.Sleep(time.Microsecond)
	wg.Wait()


}
